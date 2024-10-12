package TencentCos

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gsync/logger"
	"path"
)

type FileManager struct {
	client *cos.Client
}

func NewFileManager(bucketName string) *FileManager {
	bucketUrl, ok := FindBucketUrlByName(bucketName)
	if !ok {
		logger.Error("无法创建新的FileUploader.client,无法找到bucket[ %s ]", bucketName)
		return nil
	} else {
		logger.Info("存储桶连接成功[ %s ]", bucketUrl)
	}
	return &FileManager{
		client: NewClient(bucketUrl),
	}
}

func (f *FileManager) DefaultUpload(remoteDstDir, localSrc string) *cos.CompleteMultipartUploadResult {
	opt := &cos.MultiUploadOptions{
		// 开始分块文件大小
		PartSize: 32,
		// 上传线程池大小
		ThreadPoolSize: 16,
		// 是否开启断点续传
		CheckPoint: true,
	}
	return f.upload(remoteDstDir, localSrc, opt)
}

func (f *FileManager) upload(remoteDstDir, localSrc string, opt *cos.MultiUploadOptions) *cos.CompleteMultipartUploadResult {
	if !FolderFormatCheck(remoteDstDir) {
		logger.Error("文件上传失败,文件夹格式异常[ %s ]", remoteDstDir)
		return nil
	}
	remoteFilePath := path.Join(remoteDstDir, path.Base(localSrc))
	res, _, err := f.client.Object.Upload(
		context.Background(), remoteFilePath, localSrc, nil,
	)
	if err != nil {
		logger.Error("上传文件失败[ %s ] %s", remoteDstDir, err)
	} else {
		logger.Info("上传文件成功[ %s -> %s ]", localSrc, remoteFilePath)
	}
	return res
}

func (f *FileManager) DefaultDownload(remoteSrc, localDst string) {
	opt := &cos.MultiDownloadOptions{
		PartSize:       32,
		ThreadPoolSize: 16,
		CheckPoint:     true,
	}
	f.download(remoteSrc, localDst, opt)
}

func (f *FileManager) download(remoteSrc, localDst string, opt *cos.MultiDownloadOptions) {
	_, err := f.client.Object.Download(
		context.Background(), remoteSrc, localDst, opt,
	)
	if err != nil {
		logger.Error("文件下载失败[ %s ]", remoteSrc, err)
	} else {
		logger.Info("文件下载成功[ %s -> %s ]", remoteSrc, localDst)
	}

}

func (f *FileManager) List() (*cos.BucketGetResult, bool) {
	return f.list(nil)
}

func (f *FileManager) AutoList(remoteDstPath string) []*cos.BucketGetResult {
	var (
		bucketResultList []*cos.BucketGetResult
		hasNext          bool
		ret              *cos.BucketGetResult
	)

	if remoteDstPath == "" {
		ret, hasNext = f.list(nil)
	} else {
		ret, hasNext = f.ListDir(remoteDstPath)
	}

	bucketResultList = append(bucketResultList, ret)
	if !hasNext {
		return bucketResultList
	}
	for hasNext {
		ret, hasNext = f.ListNext(ret)
		bucketResultList = append(bucketResultList, ret)
	}
	return bucketResultList
}

func (f *FileManager) ListDir(remoteDstPath string) (*cos.BucketGetResult, bool) {
	if !FolderFormatCheck(remoteDstPath) {
		logger.Error("列出对象失败,文件夹格式异常[ %s ]", remoteDstPath)
		return nil, false
	}
	opt := &cos.BucketGetOptions{
		Prefix:    remoteDstPath, // prefix 表示要查询的文件夹
		Delimiter: "/",           // deliver 表示分隔符, 设置为/表示列出当前目录下的 object, 设置为空表示列出所有的 object
		MaxKeys:   1000,          // 设置最大遍历出多少个对象, 一次 list object 最大支持1000
	}
	return f.list(opt)
}

func (f *FileManager) ListNext(bucketResult *cos.BucketGetResult) (*cos.BucketGetResult, bool) {
	logger.Info("检测到对象列出未完整,正在继续,当前 Marker[ %s ]", bucketResult.NextMarker)
	opt := &cos.BucketGetOptions{
		Prefix:       bucketResult.Prefix,       // prefix 表示要查询的文件夹
		Delimiter:    bucketResult.Delimiter,    // deliver 表示分隔符, 设置为/表示列出当前目录下的 object, 设置为空表示列出所有的 object
		EncodingType: bucketResult.EncodingType, // 默认不编码，规定返回值的编码方式，可选值：url
		Marker:       bucketResult.NextMarker,   // 默认以 UTF-8 二进制顺序列出条目，标记返回 objects 的 list 的起点位置
		MaxKeys:      bucketResult.MaxKeys,      // 设置最大遍历出多少个对象, 一次 list object 最大支持1000
	}
	return f.list(opt)
}

func (f *FileManager) list(opt *cos.BucketGetOptions) (*cos.BucketGetResult, bool) {
	bucketResult, _, err := f.client.Bucket.Get(context.Background(), opt)
	if err != nil {
		logger.Error("列出对象失败[ %s ] %s", err)
		return nil, false
	}
	return bucketResult, bucketResult.IsTruncated
}

func (f *FileManager) DeleteFile(filePath string) {
	if _, err := f.client.Object.Delete(context.Background(), filePath); err != nil {
		logger.Error("文件删除失败[ %s ] %s", filePath, err)
	} else {
		logger.Info("文件删除成功[ %s ]", filePath)
	}
}
