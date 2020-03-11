<!DOCTYPE html>
<html lang="{{ .lang }}">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
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

<form action="/fix" method="post">
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
<form action="/fix/cover" method="post" enctype="multipart/form-data">
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
</html>