<!DOCTYPE html>
<html lang="{{ .lang }}">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>a {
            TEXT-DECORATION: none
        }</style>
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
<h1 id="title"></h1>
<table>
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
{{/*{{range .Favour}} <p> {{.}} </p> {{end}}*/}}
<div id="favour"></div>
<h4 onclick="add()" id="add_click">Add</h4>
<form action="/add/favour" id="add" method="post" style="display: none">
    <h4>Add</h4>
    <label for="id"></label><input type="number" name="id" id="id" value="{{ .Id }}" style="display: none">
    <table>
        <tr>
            <td><label for="page">page </label></td>
            <td><input type="number" name="page" id="page" placeholder="Page" required></td>
        </tr>
        <tr>
            <td><label for="contents">content </label></td>
            <td><textarea class="comments" name="contents" id="contents" rows="4" cols="50" required></textarea></td>
        </tr>
        <tr>
            <td><label for="comment">comment </label></td>
            <td><textarea class="comments" name="comment" id="comment" rows="4" cols="50"></textarea></td>
        </tr>
    </table>
    <button type="submit">Add</button>
</form>
<h4 onclick="finished()" id="finished_click">Finished</h4>
<form action="/set/time" id="finished" method="post" style="display: none">
    <h4>Finished</h4>
    <label for="ids"></label><input type="number" name="ids" id="ids" value="{{ .Id }}" style="display: none">
    <table>
        <tr>
            <td><label for="typ">Type </label></td>
            <td><select name="typ" id="typ">
                    <option value="end" selected>End</option>
                    <option value="start">Start</option>
                </select></td>
        <tr>
            <td><label for="id">Id </label></td>
            <td><input type="number" name="id" id="id" placeholder="Id" value="{{ .ReadCnt }}" required></td>
        </tr>
        <tr>
            <td><label for="time">End reading </label></td>
            <td><input type="date" name="time" id="time" placeholder="End reading" value="{{ .Time }}" required></td>
        </tr>
    </table>
    <button type="submit">Set</button>
</form>
<form action="/set/start/read" id="start" method="post" style="display: none">
    <h4>Start</h4>
    <label for="ids"></label><input type="number" name="ids" id="ids" value="{{ .Id }}" style="display: none">
    <table>
        <tr>
            <td><label for="time">Start reading </label></td>
            <td><input type="date" name="time" id="time" placeholder="Start reading" value="{{ .Time }}" required></td>
        </tr>
    </table>
    <button type="submit">Start</button>
</form>
<h4 id="fix_click" style="display: block" onclick="fix()">Fix</h4>
<form action="/fix" method="post" id="fix" style="display: none">
    <h4>Fix</h4>
    <label for="ids"></label><input type="number" name="ids" id="ids" value="{{ .Id }}" style="display: none">
    <table>
        <tr>
            <td><label for="book">Book </label></td>
            <td><input type="text" name="book" id="book" placeholder="Book" value="{{ .FixBook }}"></td>
        </tr>
        <tr>
            <td><label for="author">Author </label></td>
            <td><input type="text" name="author" id="author" placeholder="Author" value="{{ .FixAuthor }}"></td>
        </tr>
        <tr>
            <td><label for="translator">Translator </label></td>
            <td><input type="text" name="translator" id="translator" placeholder="Translator" value="{{ .FixTranslator }}"></td>
        </tr>
        <tr>
            <td><label for="publisher">Publisher </label></td>
            <td><input type="text" name="publisher" id="publisher" placeholder="Publisher" value="{{ .FixPublisher }}"></td>
        </tr>
        <tr>
            <td><label for="tag">Tag </label></td>
            <td><input type="text" name="tag" id="tag" placeholder="Tag" value="{{ .FixTag }}"></td>
        </tr>
    </table>
    <button type="submit">Fix</button>
</form>
<form action="/fix/cover" method="post" enctype="multipart/form-data" id="fix_cover"  style="display: none">
    <h4>Fix Cover</h4>
    <label for="ids"></label><input type="number" name="ids" id="ids" value="{{ .Id }}" style="display: none">
    <table>
        <tr>
            <td><label for="cover">Cover </label></td>
            <td><input type="file" name="cover" id="cover" placeholder="Cover"></td>
        </tr>
    </table>
    <button type="submit">Fix Cover</button>
</form>
</body>

<script type="text/javascript">
    let books = JSON.parse({{ .Books }});
    console.log({{.Books}});

    document.getElementById("title").innerText = books.book;
    document.getElementById("author").innerText = books.author;
    document.getElementById("translator").innerText = books.translator;
    document.getElementById("publisher").innerText = books.publisher;
    document.getElementById("cover").innerHTML = `<img src="/cover/` + books.cover + `" width="70%" />`;

    let reading = "<strong>|</strong>";
    let tag = "<strong>|</strong>";
    let favour = "<br /> <strong>>-------<</strong> <br /><br />";

    for (let i = 0; i < books.reading.length; i++) {
        reading += " " + books.reading[i].start_time + " <strong>to</strong> " + books.reading[i].end_time + " <strong>|</strong>";
    }
    for (let i = 0; i < books.tag.length; i++) {
        tag += " " + books.tag[i] + "<strong> |</strong>";
    }
    for (let i = 0; i < books.favour.length; i++) {
        favour += `<div><a><strong>[ </strong>` + books.favour[i].page + `<strong> ]</strong></a><a><strong>[ </strong>` + books.favour[i].time + `<strong> ]</strong></a><strong><p>` + books.favour[i].content + `</p></strong><a><strong>[ </strong>` + books.favour[i].comment + `<strong> ]</strong></a></div><br /><strong>>-------<</strong> <br /><br />`;
    }

    document.getElementById("reading").innerHTML = reading;
    document.getElementById("tag").innerHTML = tag;
    document.getElementById("favour").innerHTML = favour;

    function add() {
        document.getElementById("add").style.display = "block";
        document.getElementById("add_click").style.display = "none";
    }
    function finished() {
        document.getElementById("finished").style.display = "block";
        document.getElementById("start").style.display = "block";
        document.getElementById("finished_click").style.display = "none";
    }
    function fix() {
        document.getElementById("fix").style.display = "block";
        document.getElementById("fix_cover").style.display = "block";
        document.getElementById("fix_click").style.display = "none";

    }
</script>
</html>