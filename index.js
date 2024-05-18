// index.js

const express = require('express');
const fetch = require('node-fetch');
const { createProxyMiddleware } = require('http-proxy-middleware');
const app = express();
const port = 3000;

// Middleware para servir archivos estáticos
app.use(express.static('public'));

// Middleware para redirigir las solicitudes al servidor de API en el puerto 8080
app.use('/api', createProxyMiddleware({ target: 'http://localhost:8080', changeOrigin: true }));

// Página principal
app.get('/', (req, res) => {
    res.sendFile(__dirname + '/public/index.html');
});

app.listen(port, () => {
    console.log(`Frontend server listening on port ${port}`);
});
