import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../utils/theme_provider.dart';

class ThemeSwitcherButton extends StatelessWidget {
  const ThemeSwitcherButton({super.key});

  @override
  Widget build(BuildContext context) {
    final themeProvider = Provider.of<ThemeProvider>(context);
    final isDark = themeProvider.themeMode == ThemeMode.dark;
    return IconButton(
      icon: Icon(isDark ? Icons.wb_sunny : Icons.nightlight_round),
      tooltip: isDark ? 'Switch to Light Mode' : 'Switch to Dark Mode',
      onPressed: () => themeProvider.toggleTheme(),
    );
  }
} 