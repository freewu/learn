// 文件操作示例
const fs = require("fs");

// 创建一个文件
fs.open('./test','w',(err,fd) => {
    if(err) throw err;
    // 写入内容
    fs.write(fd,"dsfdsfds",(err) => {
        if(err) throw err;
    });

    // 关闭文件句柄
    fs.close(fd,(err) => {
        if(err) throw err;

        // 读取
        fs.open("./test","r",(err,fd) => {
            if(err) throw err;

            var buf = new Buffer.alloc(1024);
            fs.read(fd,buf,0,buf.length,0,(err,bytes) => {
                if(err) throw err;

                //console.log(bytes);
                // 仅输出读取的字节
                if(bytes > 0) {
                    console.log(buf.slice(0, bytes).toString());
                }
            });
            //console.log(buf);

            fs.close(fd,(err) => {
                if(err) throw err;
            });
        });
    });
});

// sleep 2 seconds
setTimeout(() => {
    
    // 直接写入
    fs.writeFile("./test","\nfsdfsdf",(err) => {
        if(err) throw err;

        // 直接全量读取
        fs.readFile('./test', function (err, data) {
            if (err) {
            return console.error(err);
            }
            console.log("异步读取文件数据: " + data.toString());
        });
    });


    setTimeout(() => {
        // 文件属性判断
        fs.stat("./test",(err,stat) => {
            if(err) throw err;
            console.log("isFile:" + stat.isFile());
            console.log("isDirectory:" + stat.isDirectory());
            console.log("mtime:" + stat.mtime);
            console.log("mode:" + stat.mode);
        });

        // 文件改名
        fs.rename('./test', './test1', (err) => {
            if (err) throw err;
            console.log('renamed complete');

            // 删除文件
            fs.unlink("./test1",(err) => {
                if(err) throw err;
                console.log("successfully delete file");
            });

            try {
                fs.unlinkSync("./test1")
                console.log("successfully delete file");
            } catch(err) {
                //console.log(err)
            }
        });
    },1000);

}, 1000);