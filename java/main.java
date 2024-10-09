import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import org.json.JSONArray;
import org.json.JSONObject;

public class Main {
    public static void main(String[] args) {
        String apiUrl = "https://apiv.forcescripts.com/DEMO";

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(apiUrl))
                .build();

        try {
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
            JSONObject data = new JSONObject(response.body());

            JSONArray matches = data.getJSONArray("data");
            for (int i = 0; i < matches.length(); i++) {
                JSONObject match = matches.getJSONObject(i);
                System.out.println("Takımlar: " + match.getString("home_team") + " vs " + match.getString("away_team"));
                System.out.println("Skor: " + match.getJSONObject("score").getJSONObject("homeTeam").getInt("current")
                        + " - " + match.getJSONObject("score").getJSONObject("awayTeam").getInt("current"));
                System.out.println("Dakika: " + match.getJSONObject("score").getInt("minute"));
                System.out.println("Başlama Saati: " + match.getLong("date"));
                System.out.println();
            }
        } catch (Exception e) {
            System.out.println("Hata: " + e.getMessage());
        }
    }
}
