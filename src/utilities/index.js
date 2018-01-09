export const capitalize = str => [str[0].toUpperCase(), str.slice(1)].join('');

export const get404HTML = () => `
<html>
<head>
  <link rel="stylesheet" type="text/css" href="http://fonts.googleapis.com/css?family=Bungee">
  <style>
    html, body, .container {
        height: 100%;
        margin: 0; 
        overflow: hidden;
    }
    .container {
        display: flex;
        align-items: center;
        justify-content: center;
        font-family: Bungee;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>404: Not Found</h1>
  </div>
</body>
</html>
`;
