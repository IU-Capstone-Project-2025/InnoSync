import 'package:flutter/material.dart';
import 'colors.dart';
import 'text_styles.dart';
import 'dimensions.dart';

class AppTheme {
  // Button Styles
  static ButtonStyle primaryButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: AppColors.primary,
    foregroundColor: AppColors.textOnPrimary,
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
    ),
    minimumSize: const Size(double.infinity, AppDimensions.buttonHeight),
  );

  static ButtonStyle secondaryButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: AppColors.error,
    foregroundColor: AppColors.textOnPrimary,
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
    ),
    minimumSize: const Size(double.infinity, AppDimensions.buttonHeight),
  );

  // Input Field Styles
  static InputDecoration inputDecoration({
    required IconData icon,
    required String hint,
    Widget? suffixIcon,
  }) {
    return InputDecoration(
      prefixIcon: Icon(icon, color: AppColors.primary),
      suffixIcon: suffixIcon,
      hintText: hint,
      border: OutlineInputBorder(
        borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
      ),
      contentPadding: const EdgeInsets.symmetric(
        vertical: AppDimensions.paddingMd,
      ),
      errorStyle: AppTextStyles.errorText,
    );
  }

  // Container Styles
  static BoxDecoration tabSelectorDecoration = BoxDecoration(
    border: Border.all(color: AppColors.primaryDark),
    borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
  );

  static BoxDecoration tabIndicatorDecoration = BoxDecoration(
    color: AppColors.primary,
    borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
  );

  // Card decoration
  static BoxDecoration cardDecoration = BoxDecoration(
    color: AppColors.cardBackground,
    borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
    boxShadow: [
      BoxShadow(
        color: AppColors.shadow,
        blurRadius: AppDimensions.shadowBlur,
        offset: const Offset(0, AppDimensions.shadowOffset),
      ),
    ],
  );

  // Light ThemeData
  static final ThemeData lightTheme = ThemeData(
    brightness: Brightness.light,
    primaryColor: AppColors.primary,
    scaffoldBackgroundColor: AppColors.background,
    cardColor: AppColors.cardBackground,
    appBarTheme: const AppBarTheme(
      backgroundColor: AppColors.background,
      foregroundColor: AppColors.textPrimary,
      elevation: 0,
      iconTheme: IconThemeData(color: AppColors.textPrimary),
      titleTextStyle: TextStyle(
        color: AppColors.textPrimary,
        fontSize: 20,
        fontWeight: FontWeight.bold,
      ),
    ),
    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: AppColors.textPrimary),
      bodyMedium: TextStyle(color: AppColors.textSecondary),
    ),
    colorScheme: ColorScheme.light(
      primary: AppColors.primary,
      surface: AppColors.surface,
      error: AppColors.error,
      onPrimary: AppColors.textOnPrimary,
      onSurface: AppColors.textPrimary,
      onError: AppColors.textOnPrimary,
    ),
    tabBarTheme: TabBarThemeData(
      indicator: BoxDecoration(
        color: AppColors.primary,
        borderRadius: BorderRadius.all(Radius.circular(AppDimensions.radiusLg)),
        border: Border.all(color: AppColors.primary, width: 2),
      ),
      dividerColor: Colors.transparent,
    ),
  );

  // Dark ThemeData
  static final ThemeData darkTheme = ThemeData(
    brightness: Brightness.dark,
    primaryColor: AppColors.primaryDarkTheme,
    scaffoldBackgroundColor: AppColors.backgroundDark,
    cardColor: AppColors.cardBackgroundDark,
    appBarTheme: const AppBarTheme(
      backgroundColor: AppColors.backgroundDark,
      foregroundColor: AppColors.textPrimaryDark,
      elevation: 0,
      iconTheme: IconThemeData(color: AppColors.textPrimaryDark),
      titleTextStyle: TextStyle(
        color: AppColors.textPrimaryDark,
        fontSize: 20,
        fontWeight: FontWeight.bold,
      ),
    ),
    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: AppColors.textPrimaryDark),
      bodyMedium: TextStyle(color: AppColors.textSecondaryDark),
    ),
    colorScheme: ColorScheme.dark(
      primary: AppColors.primaryDarkTheme,
      surface: AppColors.surfaceDark,
      error: AppColors.error,
      onPrimary: AppColors.textOnPrimaryDark,
      onSurface: AppColors.textPrimaryDark,
      onError: AppColors.textOnPrimaryDark,
    ),
    tabBarTheme: TabBarThemeData(
      indicator: BoxDecoration(
        color: AppColors.primaryDarkTheme,
        borderRadius: BorderRadius.all(Radius.circular(AppDimensions.radiusLg)),
        border: Border.all(color: AppColors.primaryDarkTheme, width: 2),
      ),
      dividerColor: Colors.transparent,
    ),
  );
}
