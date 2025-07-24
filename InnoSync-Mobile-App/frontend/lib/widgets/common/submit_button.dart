import 'package:flutter/material.dart';
import '../../theme/app_theme.dart';
import '../common/widgets.dart';

class SubmitButton extends StatelessWidget {
  final String text;
  final VoidCallback? onPressed;
  final bool isLoading;

  const SubmitButton({
    super.key,
    required this.text,
    this.onPressed,
    required this.isLoading,
  });

  @override
  Widget build(BuildContext context) {
    final colorScheme = Theme.of(context).colorScheme;
    final textTheme = Theme.of(context).textTheme;
    return ElevatedButton(
      style: AppTheme.primaryButtonStyle,
      onPressed: isLoading ? null : onPressed,
      child: isLoading
          ? FixedHeightSpace(
              height: 20,
              child: FixedWidthSpace(
                width: 20,
                child: CircularProgressIndicator(
                  strokeWidth: 2,
                  valueColor: AlwaysStoppedAnimation<Color>(
                    colorScheme.onPrimary,
                  ),
                ),
              ),
            )
          : Text(
              text,
              style: textTheme.labelLarge?.copyWith(
                color: colorScheme.onPrimary,
              ),
            ),
    );
  }
}
