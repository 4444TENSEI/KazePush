$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o ./tmp/KazePush.exe ./

# 要压缩的文件和文件夹路径
$PathsToCompress = @(
    "./tmp/KazePush.exe"
    "./public"
    "./config.toml"
)
# 压缩
Compress-Archive -Path $PathsToCompress -DestinationPath "./KazePush_win_amd64_dist.zip" -Force
# 输出执行成功提示信息
Write-Output "------ Build success(windows version), final output path: [$ZIPFILE] ------"
