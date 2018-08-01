package seelog

import (
	"net/http"
)

// 展示页面
func page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/html")
	w.WriteHeader(200)
	content := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>seelog</title>
    <link rel="shortcut icon" href="http://www.eiguo.cn/favicon.ico" type="image/x-icon" />
    <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/html2canvas/0.5.0-beta4/html2canvas.min.js"></script>
    <script>

        var out = true

        function connect (){
            var ws = new WebSocket("ws://"+ window.location.host +"/ws");
            ws.onmessage = function(e) {
                if (out){
                     $('#log').append("<p style='color: white'>"+ e.data +"</p>").scrollTop($('#log')[0].scrollHeight)
                }
            };
            ws.onclose = function () {
                $('#status').css("background-color","red").text("链接断开")
                reConnect()
            }
            ws.onopen = function () {
                $('#status').css("background-color","chartreuse").text("连接成功")
            }
            ws.onerror = function (e) {
                $('#status').css("background-color","red").text("链接断开")
            }
        }

        function reConnect(){
            setTimeout(function(){
                connect();
            },1000);
        }
        connect()


        $(function () {
            //  暂停
            $('#pause').click(function () {
                out = !out
                if (out){
                    $(this).text('暂停').css("background-color","")
                }else{
                    $(this).text('已暂停').css("background-color","red")
                }
            })

            // 清屏
            $('#clear').click(function () {
                $('#log').empty()
            })

            // 截屏
            $('#cut').click(function () {
                printPhoto("log")
            })
        })


        // 截屏
        function printPhoto(tab){
            html2canvas(document.querySelector("#"+tab)).then(canvas => {

                // 图片导出为 png 格式
                var type = 'png';
                var imgData = canvas.toDataURL(type);
                var _fixType = function(type) {
                    type = type.toLowerCase().replace(/jpg/i, 'jpeg');
                    var r = type.match(/png|jpeg|bmp|gif/)[0];
                    return 'image/' + r;
                };

                // 加工image data，替换mime type
                imgData = imgData.replace(_fixType(type),'image/octet-stream');

                //console.log(imgData);

                var saveFile = function(data, filename){
                    var save_link = document.createElementNS('http://www.w3.org/1999/xhtml', 'a');
                    save_link.href = data;
                    save_link.download = filename;

                    var event = document.createEvent('MouseEvents');
                    event.initMouseEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
                    save_link.dispatchEvent(event);
                };

                // 下载后的文件名
                var filename = 'seelog'+ '.' + type;
                // download
                saveFile(imgData,filename);

            });
        }

    </script>

</head>
<body>

<header>
    <h2 id="title">实时查看日志信息 &nbsp;<button id="status" style="background-color: darkorange">正在连接...</button></h2>
    <div class="tool">
        <button id="pause">暂停</button>
        <button id="clear">清屏</button>
        <button id="cut">截图</button>
    </div>
</header>
<div id="log"></div>
</body>

<style>
    body {
        margin-left: 2%
    }
    #title {

    }
    #log {
        width:96%;
        height: 800px;
        background-color:#181818;
        border: 1px #ccc solid;
        overflow-y: scroll;
        margin-top: 10px;
        padding-left: 12px;
        float: left;
    }
    .tool button {
        height: 30px;
        width: 100px;
        font-size: medium;
    }
</style>
</html>`
	w.Write([]byte(content))
	w.Header()
}