# 要压缩的文件和文件夹路径
$PathsToCompress = @(
    "internal"
    "public"
    "script"
    ".gitignore"
    "config.toml"
    "config.toml.example"
    "go.mod"
    "go.sum"
    "main.go"
)

# 压缩
Compress-Archive -Path $PathsToCompress -DestinationPath "KazePush_Source.zip" -Force
# 输出执行成功提示信息
Write-Output "------The source code has been backed up successfully, final output path: [$ZIPFILE]------"