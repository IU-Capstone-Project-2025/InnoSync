import 'package:flutter/material.dart';
import '../../theme/colors.dart';

class RememberMeCheckbox extends StatelessWidget {
  final bool value;
  final ValueChanged<bool?> onChanged;
  final VoidCallback? onForgotPassword;

  const RememberMeCheckbox({
    super.key,
    required this.value,
    required this.onChanged,
    this.onForgotPassword,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Checkbox(value: value, onChanged: onChanged),
        const Text("Remember me"),
        const Spacer(),
        if (onForgotPassword != null)
          TextButton(
            onPressed: onForgotPassword,
            child: Text(
              "Forgot Password ?",
              style: TextStyle(color: AppColors.info),
            ),
          ),
      ],
    );
  }
}
