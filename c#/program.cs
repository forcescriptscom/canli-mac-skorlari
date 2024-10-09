using System;
using System.Net.Http;
using System.Threading.Tasks;
using Newtonsoft.Json.Linq;

class Program
{
    static async Task Main(string[] args)
    {
        string apiUrl = "https://apiv.forcescripts.com/DEMO";

        using HttpClient client = new HttpClient();
        
        try
        {
            var response = await client.GetStringAsync(apiUrl);
            var data = JObject.Parse(response);

            foreach (var match in data["data"])
            {
                Console.WriteLine($"Takımlar: {match["home_team"]} vs {match["away_team"]}");
                Console.WriteLine($"Skor: {match["score"]["homeTeam"]["current"]} - {match["score"]["awayTeam"]["current"]}");
                Console.WriteLine($"Dakika: {match["score"]["minute"]}");
                Console.WriteLine($"Başlama Saati: {DateTimeOffset.FromUnixTimeSeconds((long)match["date"])}");
                Console.WriteLine();
            }
        }
        catch (Exception e)
        {
            Console.WriteLine($"Hata: {e.Message}");
        }
    }
}
