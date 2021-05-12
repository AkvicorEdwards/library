package handler

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	Fprint(w, `
<!DOCTYPE html>
<html lang="en" data-color-mode="auto" data-light-theme="light" data-dark-theme="dark">
<head>
    <meta charset="UTF-8">
    <title>Akvicor's Library</title>
    <style>a{TEXT-DECORATION:none}</style>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<style>
    #content {
        width: auto;
        margin-left: 10%;
        margin-right: 10%;
        margin-top: 2%;
    }
</style>
<body id="content">
<table>
    <tr><td><strong><a href="/page/book/add">Add book</a></strong></td></tr>
    <tr><td><strong><a href="/page/shelf">Book Shelf</a></strong><br/></td></tr>
    <tr><td><strong><a href="/page/likes">Likes</a></strong><br/></td></tr>
</table>
</body>
</html>
`)
}

func pageBookAdd(w http.ResponseWriter, r *http.Request) {
	Fprint(w, `<!DOCTYPE html>
<html lang="zh-cn" data-color-mode="auto" data-light-theme="light" data-dark-theme="dark">
<head>
    <meta charset="UTF-8">
    <title>Add Book</title>
    <style>a{TEXT-DECORATION:none}</style>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>
        #content {
            width: auto;
            margin-left: 10%;
            margin-right: 10%;
            margin-top: 2%;
        }
    </style>
</head>

<body id="content">
<form action="/api/book/add" method="post">
    <table>
        <tr>
            <td><label for="title">Title</label></td>
            <td><input type="text" name="title" id="title"></td>
        </tr>
        <tr>
            <td><label for="title-origin">Title Origin</label></td>
            <td><input type="text" name="title-origin" id="title-origin"></td>
        </tr>
        <tr>
            <td><label for="author">Author *M</label></td>
            <td><input type="text" name="author" id="author"></td>
        </tr>
        <tr>
            <td><label for="translator">Translator *M</label></td>
            <td><input type="text" name="translator" id="translator"></td>
        </tr>
        <tr>
            <td><label for="publisher">Publisher</label></td>
            <td><input type="text" name="publisher" id="publisher"></td>
        </tr>
        <tr>
            <td><label for="cover">Cover</label></td>
            <td><input type="text" name="cover" id="cover"></td>
        </tr>
        <tr>
            <td><label for="tag">Tag *M</label></td>
            <td><input type="text" name="tag" id="tag"></td>
        </tr>
    </table>
    <button type="submit">Submit</button>
    <br />
    <h4 id="fix_click" style="display: block" onclick="location='/'">Home</h4>
</form>
</body>

</html>
`)
}

func pageShelf(w http.ResponseWriter, r *http.Request) {
	Fprint(w, `
<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <title>Book Shelf</title>
    <style>a{TEXT-DECORATION:none}</style>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
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

<body id="content"></body>

<script type="text/javascript">

    function getContent(url) {
        let httpRequest = new XMLHttpRequest();
        httpRequest.open('GET', url, true);
        httpRequest.send();
        httpRequest.onreadystatechange = function () {
            if (httpRequest.readyState === 4 && httpRequest.status === 200) {
                let shelf = JSON.parse(httpRequest.responseText);
                if (shelf.status === 0) {
                    let shelf2 = shelf.data
                    let html = '<ul id="filelist">';
                    for(let i = 0; i < shelf2.length; i++) {
                        html += '<li><a href="/page/book?b='+shelf2[i].id+'">'+shelf2[i].title+'</a></li>';
                    }
                    html += '</ul><h4 id="fix_click" style="display: block" onclick="location=\'/\'">Home</h4>';
                    document.getElementById("content").innerHTML = html;
                }
            }
        };
    }
    getContent("/api/shelf")

</script>
</html>
`)
}

func pageBook(w http.ResponseWriter, r *http.Request) {
	bid := Int64(r.FormValue("b"))
	Fprint(w, `
<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <title>Likes</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>a {
        TEXT-DECORATION: none
    }</style>
    <style>
        #body {
            width: auto;
        }
        .comments {
            /*width: 100%;*/
            overflow: auto;
            word-break: break-all;
        }
        .content {
            /*width: 100%;*/
            overflow: auto;
            word-break: break-all;
        }
    </style>
</head>

<body id="body">
<h1 id="title"></h1>
<table>
    <tr>
        <td>Title Origin:</td>
        <td id="title-origin"></td>
    </tr>
    <tr>
        <td>Author:</td>
        <td id="author"></td>
    </tr>
    <tr>
        <td>Translator:</td>
        <td id="translator"></td>
    </tr>
    <tr>
        <td>Publisher:</td>
        <td id="publisher"></td>
    </tr>
    <tr>
        <td>Tag:</td>
        <td id="tag"></td>
    </tr>
    <tr>
        <td>Reading:</td>
        <td id="reading"></td>
    </tr>
    <tr>
        <td>Cover:</td>
        <td id="cover"></td>
    </tr>
</table>

<div id="favour"></div>
<h4 onclick="add()" id="add_click">Add</h4>
<form action="/api/book/like/add" id="add" method="post" style="display: none">
	<label for="bid"></label><input type="number" name="bid" id="bid" value="`+fmt.Sprint(bid)+`" style="display: none">
    <h4>Add</h4>
    <table>
        <tr>
            <td><label for="page">Page</label></td>
            <td><input type="text" name="page" id="page"></td>
        </tr>
        <tr>
            <td><label for="content">Content</label></td>
            <td><textarea class="comments" name="content" id="content" rows="4" cols="50" required></textarea></td>
        </tr>
        <tr>
            <td><label for="comment">Comment</label></td>
            <td><textarea class="comments" name="comment" id="comment" rows="4" cols="50"></textarea></td>
        </tr>
    </table>
    <button type="button" onclick="postForm()">Add</button>
</form>

<h4 onclick="start()">Reading Start</h4>
<h4 onclick="finished()">Reading Finished</h4>

<h4 id="fix_click" style="display: block" onclick="fix()">Fix</h4>
<form action="/api/book/fix" method="post" id="fix" style="display: none">
    <label for="b-id"></label><input type="number" name="b" id="b-id" style="display: none">
    <table>
        <tr>
            <td><label for="b-title">Title</label></td>
            <td><input type="text" name="title" id="b-title"></td>
        </tr>
        <tr>
            <td><label for="b-title-origin">Title Origin</label></td>
            <td><input type="text" name="title-origin" id="b-title-origin"></td>
        </tr>
        <tr>
            <td><label for="b-author">Author *M</label></td>
            <td><input type="text" name="author" id="b-author"></td>
        </tr>
        <tr>
            <td><label for="b-translator">Translator *M</label></td>
            <td><input type="text" name="translator" id="b-translator"></td>
        </tr>
        <tr>
            <td><label for="b-publisher">Publisher</label></td>
            <td><input type="text" name="publisher" id="b-publisher"></td>
        </tr>
        <tr>
            <td><label for="b-cover">Cover</label></td>
            <td><input type="text" name="cover" id="b-cover"></td>
        </tr>
        <tr>
            <td><label for="b-tag">Tag *M</label></td>
            <td><input type="text" name="tag" id="b-tag"></td>
        </tr>
    </table>
    <button type="submit">Submit</button>
    <br />
    <h4 id="fix_click" style="display: block" onclick="location='/'">Home</h4>
</form>
<h4 id="fix_click" style="display: block" onclick="location='/'">Home</h4>
</body>
<script type="text/javascript">
    function getQueryVariable(variable) {
        let query = window.location.search.substring(1);
        let vars = query.split("&");
        for (let i=0;i<vars.length;i++) {
            let pair = vars[i].split("=");
            if(pair[0] == variable){return pair[1];}
        }
        return(false);
    }

    function add() {
        document.getElementById("add").style.display = "block";
        document.getElementById("add_click").style.display = "none";
    }

    function postForm() {
        document.getElementById("add").submit()
    }

    function del(id) {
        if (document.getElementById("f"+id).elements['comment'].value === '7') {
            document.getElementById("d"+id).submit()
        }else {
            alert('Delete failed')
        }
    }

    function HTMLEncode(html) {
        let temp = document.createElement("div");
        (temp.textContent != null) ? (temp.textContent = html) : (temp.innerText = html);
        let output = temp.innerHTML;
        temp = null;
        return output;
    }

	let book = ""
    let likes = ""
	let reads = ""

    function getContent(url) {
        let httpRequest = new XMLHttpRequest();
        httpRequest.open('GET', url, true);
        httpRequest.send();
        httpRequest.onreadystatechange = function () {
            if (httpRequest.readyState === 4 && httpRequest.status === 200) {
                let shelf = JSON.parse(httpRequest.responseText);
                if (shelf.status === 0) {
                    book = shelf.data
                    likes = book.likes
                    reads = book.reads
					addInfo()
					addFavour()
                }
            }
        };
    }
    function addInfo() {
        document.getElementById("title").innerText = book.title;
        document.getElementById("title-origin").innerText = book.title_origin;
        document.getElementById("author").innerText = book.author;
        document.getElementById("translator").innerText = book.translator;
        document.getElementById("publisher").innerText = book.publisher;
        document.getElementById("cover").innerHTML = "<img src=\"" + book.cover + "\" width=\"70%\" />";
        let reading = "<strong>|</strong>";
        for (let i = 0; i < reads.length; i++) {
            reading += " " + reads[i].start + " <strong>to</strong> " + reads[i].end + " <strong>|</strong>";
        }
        document.getElementById("reading").innerHTML = reading;
        let tag = "<strong>|</strong>";
        for (let i = 0; i < book.tag.length; i++) {
            tag += " " + book.tag[i] + "<strong> |</strong>";
        }
        document.getElementById("tag").innerHTML = tag;
    }
    function addFavour() {
        console.log(likes)
        let favour = "<br /> <strong>-----------------</strong> <br /><br />";
        for (let i = 0; i < likes.length; i++) {
            let cont = HTMLEncode(likes[i].content).replace(" ","&nbsp;").split(/\r?\n/);
            let comm = HTMLEncode(likes[i].comment).replace(" ","&nbsp;").split(/\r?\n/);

            favour += "<div id='"+i+"'><a><strong>[ </strong>"
                + likes[i].page +
                "<strong> ]</strong></a><a><strong>[ </strong>"
                + likes[i].time +
                "<strong> ]</strong></a><br /><br /><strong>";

            // favour += books.favour[i].content;
            for (let i = 0; i < cont.length; i++) {
                favour += "<a>"+cont[i] + "</a><br />";
            }

            favour += "</strong><br />";

            // favour += books.favour[i].comment;
            for (let i = 0; i < comm.length; i++) {
                favour += "<a>"+comm[i] + "</a><br />";
            }

            favour += " </div>"
                + "<br /><strong onclick=\"fixFavour('"+i+"')\">-----------------</strong> <br /><br />";

        }
        document.getElementById("favour").innerHTML = favour;
    }
    function fix() {
        document.getElementById("fix").style.display = "block";
        document.getElementById("b-id").value = book.id;
        document.getElementById("b-title").value = book.title;
        document.getElementById("b-title-origin").value = book.title_origin;
        document.getElementById("b-author").value = book.author;
        document.getElementById("b-translator").value = book.translator;
        document.getElementById("b-publisher").value = book.publisher;
        document.getElementById("b-cover").value = book.cover;
        document.getElementById("b-tag").value = book.tag;
    }
    function fixFavour(id) {
        let f = '<form action="/api/book/like/del" method="post" id="d'+id+'" style="display: none">';
        f += '<label for="id"></label><input type="number" name="id" id="id" value="'+likes[id].id+'" style="display: none">';
        f += '</form>';
        f += "<strong onclick=\"del('"+id+"')\"> (╯°Д°)╯ ┴─┴ </strong> <br /><br />";
        f += '<form action="/api/book/like/fix" method="post" id="f'+id+'">';
        f += '<label for="id"></label><input type="number" name="id" id="id" value="'+likes[id].id+'" style="display: none">';
        f += '<table>';
        f += '<tr><td><label for="page">Page</label></td>';
        f += '<td><input type="text" name="page" id="page" value="'+likes[id].page+'" required></td></tr>';
        f += '<tr><td><label for="content">content </label></td>';
        f += '<td><textarea class="comments" name="content" id="content" rows="4" cols="50" required>'+ likes[id].content +'</textarea></td></tr>';
        f += '<tr><td><label for="comment">comment </label></td>';
        f += '<td><textarea class="comments" name="comment" id="comment" rows="4" cols="50">'+ likes[id].comment +'</textarea></td></tr>';
        f += '</table>';
        f += '</form>';
        f += "<br /><strong onclick=\"document.getElementById('f"+id+"').submit()\"> ┬─┬ ノ( ' - 'ノ) </strong><br />";
        document.getElementById(id).innerHTML = f
    }

    function start() {
        let res = confirm("Reading Start?")
        if (res) {
            let httpRequest = new XMLHttpRequest();
            httpRequest.open('GET', "/api/book/read/start?bid="+book.id, true);
            httpRequest.send();
            httpRequest.onreadystatechange = function () {
                if (httpRequest.readyState === 4 && httpRequest.status === 200) {
                    let shelf = JSON.parse(httpRequest.responseText);
                    if(shelf.status==0) {
                        alert("OK")
                    } else {
                        alert("ER")
                    }
                }
            };
        }
    }
    function finished() {
        let res = confirm("Reading Finished?")
        if (res) {
            let httpRequest = new XMLHttpRequest();
            httpRequest.open('GET', "/api/book/read/finish?bid="+book.id, true);
            httpRequest.send();
            httpRequest.onreadystatechange = function () {
                if (httpRequest.readyState === 4 && httpRequest.status === 200) {
                    let shelf = JSON.parse(httpRequest.responseText);
                    if(shelf.status==0) {
                        alert("OK")
                    } else {
                        alert("ER")
                    }
                }
            };
        }
    }

    getContent("/api/book?b="+getQueryVariable('b'))
</script>

</html>
`)
}

func pageLikes(w http.ResponseWriter, r *http.Request) {
	Fprint(w, `
<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <title>Likes</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>a {
            TEXT-DECORATION: none
        }</style>
    <style>
        #content {
            width: auto;
        }
        #comments {
            width: auto;
        }
        .comments {
            /*width: 100%;*/
            overflow: auto;
            word-break: break-all;
        }
        .content {
            /*width: 100%;*/
            overflow: auto;
            word-break: break-all;
        }
    </style>
</head>

<body id="content">
<h1 id="title"></h1>

<div id="favour"></div>
<h4 onclick="add()" id="add_click">Add</h4>
<form action="/api/like/add" id="add" method="post" style="display: none">
    <h4>Add</h4>
    <table>
        <tr>
            <td><label for="origin">Origin</label></td>
            <td><input type="text" name="origin" id="origin"></td>
        </tr>
        <tr>
            <td><label for="content">Content</label></td>
            <td><textarea class="comments" name="content" id="content" rows="4" cols="50" required></textarea></td>
        </tr>
        <tr>
            <td><label for="comment">Comment</label></td>
            <td><textarea class="comments" name="comment" id="comment" rows="4" cols="50"></textarea></td>
        </tr>
    </table>
    <button type="button" onclick="postForm()">Add</button>
</form>

<h4 id="fix_click" style="display: block" onclick="location='/'">Home</h4>
</body>
<script type="text/javascript">
    function add() {
        document.getElementById("add").style.display = "block";
        document.getElementById("add_click").style.display = "none";
    }

    function postForm() {
        document.getElementById("add").submit()
    }

    function del(id) {
        if (document.getElementById("f"+id).elements['comment'].value === '7') {
            document.getElementById("d"+id).submit()
        }else {
            alert('Delete failed')
        }
    }

    function HTMLEncode(html) {
        let temp = document.createElement("div");
        (temp.textContent != null) ? (temp.textContent = html) : (temp.innerText = html);
        let output = temp.innerHTML;
        temp = null;
        return output;
    }

    let likes = ""

    function getContent(url) {
        let httpRequest = new XMLHttpRequest();
        httpRequest.open('GET', url, true);
        httpRequest.send();
        httpRequest.onreadystatechange = function () {
            if (httpRequest.readyState === 4 && httpRequest.status === 200) {
                let shelf = JSON.parse(httpRequest.responseText);
                if (shelf.status === 0) {
                    likes = shelf.data
                    addFavour()
                }
            }
        };
    }
    function addFavour() {
        console.log(likes)
        let favour = "<br /> <strong>-----------------</strong> <br /><br />";
        for (let i = 0; i < likes.length; i++) {
            let cont = HTMLEncode(likes[i].content).replace(" ","&nbsp;").split(/\r?\n/);
            let comm = HTMLEncode(likes[i].comment).replace(" ","&nbsp;").split(/\r?\n/);

            favour += "<div id='"+i+"'><a><strong>[ </strong>"
                + likes[i].origin +
                "<strong> ]</strong></a><a><strong>[ </strong>"
                + likes[i].time +
                "<strong> ]</strong></a><br /><br /><strong>";

            // favour += books.favour[i].content;
            for (let i = 0; i < cont.length; i++) {
                favour += "<a>"+cont[i] + "</a><br />";
            }

            favour += "</strong><br />";

            // favour += books.favour[i].comment;
            for (let i = 0; i < comm.length; i++) {
                favour += "<a>"+comm[i] + "</a><br />";
            }

            favour += " </div>"
                + "<br /><strong onclick=\"fixFavour('"+i+"')\">-----------------</strong> <br /><br />";

        }
        document.getElementById("favour").innerHTML = favour;
    }
    function fixFavour(id) {
        let f = '<form action="/api/like/del" method="post" id="d'+id+'" style="display: none">';
        f += '<label for="id"></label><input type="number" name="id" id="id" value="'+likes[id].id+'" style="display: none">';
        f += '</form>';
        f += "<strong onclick=\"del('"+id+"')\"> (╯°Д°)╯ ┴─┴ </strong> <br /><br />";
        f += '<form action="/api/like/fix" method="post" id="f'+id+'">';
        f += '<label for="id"></label><input type="number" name="id" id="id" value="'+likes[id].id+'" style="display: none">';
        f += '<table>';
        f += '<tr><td><label for="origin">Origin</label></td>';
        f += '<td><input type="text" name="origin" id="origin" value="'+likes[id].origin+'" required></td></tr>';
        f += '<tr><td><label for="content">content </label></td>';
        f += '<td><textarea class="comments" name="content" id="content" rows="4" cols="50" required>'+ likes[id].content +'</textarea></td></tr>';
        f += '<tr><td><label for="comment">comment </label></td>';
        f += '<td><textarea class="comments" name="comment" id="comment" rows="4" cols="50">'+ likes[id].comment +'</textarea></td></tr>';
        f += '</table>';
        f += '</form>';
        f += "<br /><strong onclick=\"document.getElementById('f"+id+"').submit()\"> ┬─┬ ノ( ' - 'ノ) </strong><br />";
        document.getElementById(id).innerHTML = f
    }

    getContent("/api/like")
</script>

</html>
`)
}