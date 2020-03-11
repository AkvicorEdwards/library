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
    </style>
</head>

<body id="content">
<form action="/add/book" method="post" enctype="multipart/form-data">
    <table>
        <tr>
            <td><label for="book">Book </label></td>
            <td><input type="text" name="book" id="book" placeholder="Book"></td>
        </tr>
        <tr>
            <td><label for="author">Author </label></td>
            <td><input type="text" name="author" id="author" placeholder="Author"></td>
        </tr>
        <tr>
            <td><label for="translator">Translator </label></td>
            <td><input type="text" name="translator" id="translator" placeholder="Translator"></td>
        </tr>
        <tr>
            <td><label for="publisher">Publisher </label></td>
            <td><input type="text" name="publisher" id="publisher" placeholder="Publisher"></td>
        </tr>
        <tr>
            <td><label for="cover">Cover </label></td>
            <td><input type="file" name="cover" id="cover" placeholder="Cover"></td>
        </tr>
        <tr>
            <td><label for="tag">Tag </label></td>
            <td><input type="text" name="tag" id="tag" placeholder="Tag"></td>
        </tr>
        <tr>
            <td><label for="time">Start reading </label></td>
            <td><input type="date" name="time" id="time" placeholder="Start reading" value="{{ .Time }}"></td>
        </tr>
    </table>
    <button type="submit">Submit</button>
</form>
</body>

</html>