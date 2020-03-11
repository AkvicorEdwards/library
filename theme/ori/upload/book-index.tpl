<!DOCTYPE html>
<html lang="{{ .lang }}">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
    <style>a{TEXT-DECORATION:none}</style>
    <style>
        #content {
            width: auto;
            margin-left: 10%;
            margin-right: 10%;
            margin-top: 2%;
        }

        .comments {
            /*width: 100%;*/
            overflow: auto;
            word-break: break-all;
        }
    </style>

</head>

<body id="content">
</body>
<script type="text/javascript">
    let books = JSON.parse({{ .Books }});
    let urlShort = {{ .UrlShort }};
    let html = '<ul id="filelist">';
    for(let i = 0; i < books.length; i++) {
        console.log(books[i].id);
        html += '<li><a href="'+urlShort+'book/'+books[i].id+'">'+books[i].book+'</a></li>';
    }
    html += '</ul>';
    document.getElementById("content").innerHTML = html;
</script>
</html>