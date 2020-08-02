var keyword_back
function getJSON(output,flag){
    var req = new XMLHttpRequest();
    var keyword = document.getElementById("keyword").value
    if ((keyword == keyword_back)&&(flag==false)){
        return
    }
    req.onreadystatechange = function(){
        if(req.readyState == 4 && req.status == 200){
            nowserchpage = 1
            var data=req.responseText;
            var tmp = JSON.parse(data)
            console.log(tmp)
            document.getElementById(output).innerHTML = "<div>処理時間:"+tmp.Time+"</div>"+outputdiv(tmp.Data);
            keyword_back = keyword
        }
    };
    req.open("GET","./v1/isbn/"+keyword,false);
    req.send(null);
}

function outputdiv(json){
    var output = "<div clase=\"table\">"
    var title = ["ISBN","Title","作者","出版","あらすじ","その他"]
    var bodydata = [json.Isbn,json.Title,json.Writer,json.Brand,json.Synopsis,json.Ext]
    output += "<div class=\"tableRow\">"
    output += "<div class=\"tableHead\">Name</div>"
    output += "<div class=\"tableHead\">Data</div>"
    output += "</div>"
    for (var i=0;i<title.length;i++){
        output += "<div class=\"tableRow\">"
        output += "<div class=\"tableCell\">"+title[i]+"</div>"
        output += "<div class=\"tableCell\">"+bodydata[i]+"</div>"
        output += "</div>"
    }
    output += "<div class=\"tableRow\">"
    output += "<div class=\"tableCell\">"+"画像"+"</div>"
    output += "<div class=\"tableCell\">"+"<img src='"+ json.Image+"' alt='"+json.Image+"'>"+"</div>"
    output += "</div></div>"
    return output
}