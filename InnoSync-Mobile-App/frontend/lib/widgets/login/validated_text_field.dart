import 'package:flutter/material.dart';
import '../../theme/dimensions.dart';
import 'package:flutter/services.dart';

class ValidatedTextField extends StatelessWidget {
  final IconData icon;
  final String hint;
  final TextEditingController controller;
  final String? Function(String?)? validator;
  final bool obscure;
  final VoidCallback? toggleObscure;

  const ValidatedTextField({
    super.key,
    required this.icon,
    required this.hint,
    required this.controller,
    this.validator,
    this.obscure = false,
    this.toggleObscure,
  });

  @override
  Widget build(BuildContext context) {
    final colorScheme = Theme.of(context).colorScheme;
    return TextFormField(
      controller: controller,
      validator: validator,
      obscureText: obscure,
      inputFormatters: [NoControlCharsFormatter()],
      decoration: InputDecoration(
        prefixIcon: Icon(icon, color: colorScheme.primary),
        hintText: hint,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
        ),
        contentPadding: const EdgeInsets.symmetric(
          vertical: AppDimensions.paddingMd,
        ),
        suffixIcon: toggleObscure != null
            ? IconButton(
                icon: Icon(
                  obscure ? Icons.visibility_off : Icons.visibility,
                  color: colorScheme.primary,
                ),
                onPressed: toggleObscure,
              )
            : null,
      ),
    );
  }
}

class NoControlCharsFormatter extends TextInputFormatter {
  @override
  TextEditingValue formatEditUpdate(
    TextEditingValue oldValue,
    TextEditingValue newValue,
  ) {
    // Удаляем все управляющие символы (код < 32, кроме \n)
    final filtered = newValue.text.replaceAll(RegExp(r'[\x00-\x1F\x7F]'), '');
    return TextEditingValue(text: filtered, selection: newValue.selection);
  }
}
