const axios = require('axios');

const api_url = 'https://apiv.forcescripts.com/DEMO';

axios.get(api_url)
    .then(response => {
        const data = response.data;
        data.data.forEach(match => {
            console.log(`Takımlar: ${match.home_team} vs ${match.away_team}`);
            console.log(`Skor: ${match.score.homeTeam.current} - ${match.score.awayTeam.current}`);
            console.log(`Dakika: ${match.score.minute}`);
            console.log(`Başlama Saati: ${new Date(match.date * 1000).toLocaleString()}`);
            console.log();
        });
    })
    .catch(error => {
        console.log('Hata:', error.message);
    });
