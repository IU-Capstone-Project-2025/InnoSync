import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'app.dart';
import 'utils/theme_provider.dart';

void main() {
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => ThemeProvider()),
      ],
      child: const MyApp(),
    ),
  );
}
