import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart' as http;

Future<void> main() async {
  final client = http.Client();

  try {
    final request = http.Request(
      'POST',
      Uri.parse('http://127.0.0.1:9090/recipeStepsFlow'),
    );
    request.headers['Accept'] = 'text/event-stream';
    request.headers['Content-Type'] = 'application/json';
    request.body = jsonEncode({'data': 'Ramen'});

    final response = await client.send(request);

    if (response.statusCode == 200) {
      var buffer = '';

      await response.stream.transform(utf8.decoder).listen((chunk) {
        buffer += chunk;

        while (buffer.contains('\n\n')) {
          final eventEnd = buffer.indexOf('\n\n');
          final event = buffer.substring(0, eventEnd);
          buffer = buffer.substring(eventEnd + 2);

          if (event.startsWith('data: ')) {
            final jsonData = event.substring(6);
            final parsed = jsonDecode(jsonData);

            if (parsed['message'] != null) {
              print('Streaming chunk: ${parsed['message']}');
            } else if (parsed['result'] != null) {
              print('Final result: ${parsed['result']}');
            }
          }
        }
      }).asFuture();
    }
  } finally {
    client.close();
  }
}
