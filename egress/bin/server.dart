import 'dart:async';
import 'dart:io';

import 'package:shelf/shelf.dart';
import 'package:shelf/shelf_io.dart';
import 'package:shelf_router/shelf_router.dart';
import 'package:http/http.dart' as http;

// Configure routes.
final _router = Router()
  ..post('/translate', _translateHandler)
  ..get('/dictionary', _dictionaryHandler)
  ..get('/health', _healthHandler);

FutureOr<Response> _translateHandler(Request req) async {
  try {
    final headers = Map<String, String>.from(req.headers);
    headers.remove('host');
    headers.remove('content-length');

    final body = await req.read().toList();
    final bodyBytes = body.expand((x) => x).toList();

    final proxyRequest = http.Request("POST", Uri.parse('https://translate.api.cloud.yandex.net/translate/v2/translate'))
      ..headers.addAll(headers)
      ..bodyBytes = bodyBytes;

    final response = await proxyRequest.send();
    final responseBody = await response.stream.toBytes();

    return Response(
      response.statusCode,
      headers: {'content-type': response.headers['content-type'] ?? 'application/json', ...response.headers},
      body: responseBody,
    );
  } catch (e) {
    return Response(500);
  }
}

FutureOr<Response> _dictionaryHandler(Request req) async {
  try {
    final headers = Map<String, String>.from(req.headers);
    headers.remove('host');
    headers.remove('content-length');

    final body = await req.read().toList();
    final bodyBytes = body.expand((x) => x).toList();

    final proxyRequest = http.Request("POST", Uri.parse('https://dictionary.yandex.net/api/v1/dicservice/lookup'))
      ..headers.addAll(headers)
      ..bodyBytes = bodyBytes;

    final response = await proxyRequest.send();
    final responseBody = await response.stream.toBytes();

    return Response(
      response.statusCode,
      headers: {'content-type': response.headers['content-type'] ?? 'application/json', ...response.headers},
      body: responseBody,
    );
  } catch (e) {
    return Response(500);
  }
}

Response _healthHandler(Request request) {
  return Response.ok('OK');
}

void main(List<String> args) async {
  // Use any available host or container IP (usually `0.0.0.0`).
  final ip = InternetAddress.anyIPv4;

  // Configure a pipeline that logs requests.
  final handler = Pipeline().addMiddleware(logRequests()).addHandler(_router.call);

  // For running in containers, we respect the PORT environment variable.
  final port = int.parse(Platform.environment['PORT'] ?? '8080');
  final server = await serve(handler, ip, port);
  print('Server listening on port ${server.port}');
}
