<!DOCTYPE html>
<html lang="{{ .lang }}">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
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
<h1>Login</h1>
<form action="/login" method="post">
    <table>
        <tr>
            <td><label for="username">Username </label></td>
            <td><input type="text" name="username" id="username" placeholder="username"></td>
        </tr>
        <tr>
            <td><label for="password">Password </label></td>
            <td><input type="text" name="password" id="password" placeholder="password"></td>
        </tr>
    </table>
    <button type="submit">Submit</button>
</form>
</body>
</html>