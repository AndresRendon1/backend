// pages/index.js
import axios from 'axios';
import { useState } from 'react';

export default function Home() {
  const [heroName, setHeroName] = useState('');
  const [heroData, setHeroData] = useState(null);
  const [error, setError] = useState(null);

  const fetchHeroData = async () => {
    try {
      setError(null); // Reset error before fetching data
      const response = await axios.get(`http://localhost:8080/api/superhero?hero=${heroName}`);
      setHeroData(response.data);
    } catch (error) {
      console.error('Error fetching hero data:', error);
      setHeroData(null);
      setError('Error fetching hero data. Please try again.');
    }
  };

  return (
    <div className="container">
      <h1>Superhero Search</h1>
      <input
        type="text"
        value={heroName}
        onChange={(e) => setHeroName(e.target.value)}
        placeholder="Enter superhero name"
      />
      <button onClick={fetchHeroData}>Search</button>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {heroData && (
        <div>
          <h2>{heroData.name}</h2>
          <p>Full Name: {heroData.biography.fullName}</p>
          <h3>Powerstats:</h3>
          <ul>
            <li>Intelligence: {heroData.powerstats.intelligence}</li>
            <li>Strength: {heroData.powerstats.strength}</li>
            <li>Speed: {heroData.powerstats.speed}</li>
            <li>Durability: {heroData.powerstats.durability}</li>
            <li>Power: {heroData.powerstats.power}</li>
            <li>Combat: {heroData.powerstats.combat}</li>
          </ul>
          <img src={heroData.images.md} alt={heroData.name} />
        </div>
      )}
    </div>
  );
}
