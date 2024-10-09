import requests

api_url = 'https://apiv.forcescripts.com/DEMO'

try:
    response = requests.get(api_url)
    data = response.json()

    for match in data['data']:
        print(f"Takımlar: {match['home_team']} vs {match['away_team']}")
        print(f"Skor: {match['score']['homeTeam']['current']} - {match['score']['awayTeam']['current']}")
        print(f"Dakika: {match['score']['minute']}")
        print(f"Başlama Saati: {match['date']}")
        print()
except Exception as e:
    print(f'Hata: {e}')
