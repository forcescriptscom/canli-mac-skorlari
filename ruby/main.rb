require 'net/http'
require 'json'

api_url = 'https://apiv.forcescripts.com/DEMO'
uri = URI(api_url)

begin
  response = Net::HTTP.get(uri)
  data = JSON.parse(response)

  data['data'].each do |match|
    puts "Takımlar: #{match['home_team']} vs #{match['away_team']}"
    puts "Skor: #{match['score']['homeTeam']['current']} - #{match['score']['awayTeam']['current']}"
    puts "Dakika: #{match['score']['minute']}"
    puts "Başlama Saati: #{Time.at(match['date'])}"
    puts
  end
rescue StandardError => e
  puts "Hata: #{e.message}"
end
