const express = require('express');
const app = express();

// Almacenamiento en memoria de los datos de los superhéroes
const superheroes = {
  "Wolverine": {
    "biography": {
      "fullName": "John Logan"
    },
    "powerstats": {
      "intelligence": 63,
      "strength": 32,
      "speed": 50,
      "durability": 100,
      "power": 89,
      "combat": 100
    },
    "images": {
      "xs": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
      "sm": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
      "md": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
      "lg": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg"
    }
  },
  // Agrega más superhéroes aquí
};

// Endpoint para obtener los datos de un superhéroe por nombre
app.get('/api/superhero', (req, res) => {
  const heroName = req.query.hero;
  const superhero = superheroes[heroName];
  if (!superhero) {
    return res.status(404).json({ error: 'Superhero not found' });
  }
  res.json(superhero);
});

// Inicia el servidor en el puerto 8080
const PORT = 8080;
app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
