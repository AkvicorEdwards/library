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

<div id="favour"></div>
<h4 onclick="add()" id="add_click">Add</h4>
<form action="/add/favour" id="add" method="post" style="display: none">
    <h4>Add</h4>
    <label for="id"></label><input type="number" name="id" id="id" value="{{ .Id }}" style="display: none">
    <label for="page"></label><input type="text" name="page" id="page" style="display: none">
    <table>
        <tr>
            <td><label for="title">title </label></td>
            <td><input type="text" name="title" id="title" placeholder="Title"></td>
        </tr>
        <tr>
            <td><label for="author">author </label></td>
            <td><input type="text" name="author" id="author" placeholder="Author"></td>
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
    <button type="button" onclick="postForm()">Add</button>
</form>
</body>

<script type="text/javascript">
    let books = JSON.parse({{ .Books }});
    let rid = {{ .Id }};
    console.log({{.Books}});

    addFavour();

    function add() {
        document.getElementById("add").style.display = "block";
        document.getElementById("add_click").style.display = "none";
    }
    function addFavour() {
        let favour = "<br /> <strong>-----------------</strong> <br /><br />";
        for (let i = 0; i < books.favour.length; i++) {
            favour += `<div id='`+i+`'><a><strong>[ </strong>`
                + books.favour[i].page +
                `<strong> ]</strong></a><a><strong>[ </strong>`
                + books.favour[i].time +
                `<strong> ]</strong></a><strong><p>`
                + books.favour[i].content +
                `</p></strong><a><strong>[ </strong>`
                + books.favour[i].comment +
                `<strong> ]</strong></a> </div>`
                + `<br /><strong onclick="fixFavour('`+i+`')">-----------------</strong> <br /><br />`;
        }
        document.getElementById("favour").innerHTML = favour;
    }
    function fixFavour(id) {
        let f = `<form action="/del/favour" method="post" id="d`+id+`" style="display: none">`;
        f += `<label for="ids"></label><input type="number" name="ids" id="ids" value="`+id+`" style="display: none">`;
        f += `<label for="id"></label><input type="number" name="id" id="id" value="`+rid+`" style="display: none">`;
        f += `</form>`;
        f += `<strong onclick="del('`+id+`')"> (╯°Д°)╯ ┴─┴ </strong> <br /><br />`;
        f += `<form action="/fix/favour" method="post" id="f`+id+`">`;
        f += `<label for="ids"></label><input type="number" name="ids" id="ids" value="`+id+`" style="display: none">`;
        f += `<label for="id"></label><input type="number" name="id" id="id" value="`+rid+`" style="display: none">`;
        f += `<table>`;
        f += `<tr><td><label for="page">Page</label></td>`;
        f += `<td><input type="text" name="page" id="page" value="`+books.favour[id].page+`" required></td></tr>`;
        f += `<tr><td><label for="time">Time</label></td>`;
        f += `<td><input type="text" name="time" id="time" value="`+ books.favour[id].time +`" required></td></tr>`;
        f += `<tr><td><label for="contents">content </label></td>`;
        f += `<td><textarea class="comments" name="contents" id="contents" rows="4" cols="50" required>`+ books.favour[id].content +`</textarea></td></tr>`;
        f += `<tr><td><label for="comment">comment </label></td>`;
        f += `<td><textarea class="comments" name="comment" id="comment" rows="4" cols="50">`+ books.favour[id].comment +`</textarea></td></tr>`;
        f += `</table>`;
        f += `</form>`;
        f += `<br /><strong  onclick="document.getElementById('f`+id+`').submit()"> ┬─┬ ノ( ' - 'ノ) </strong><br />`;
        document.getElementById(id).innerHTML = f
    }
    function del(id) {
        if (document.getElementById("f"+id).elements['comment'].value === '7') {
            document.getElementById("d"+id).submit()
        }else {
            alert('Delete failed')
        }
    }

    function postForm() {
        document.getElementById("add").elements['page'].value = document.getElementById("add").elements['title'].value + " -- " + document.getElementById("add").elements['author'].value;
        document.getElementById("add").submit()
    }
</script>

</html>