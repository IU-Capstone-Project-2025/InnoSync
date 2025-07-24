import 'package:flutter/material.dart';
import 'validated_text_field.dart';
import 'remember_me_checkbox.dart';
import '../../widgets/common/spacing.dart';
import '../../widgets/common/submit_button.dart';
import '../../utils/validators.dart';

class LoginTab extends StatefulWidget {
  final TextEditingController emailController;
  final TextEditingController passwordController;
  final bool rememberMe;
  final Function(bool?) onRememberMeChanged;
  final VoidCallback onForgotPassword;
  final VoidCallback onLogin;

  const LoginTab({
    super.key,
    required this.emailController,
    required this.passwordController,
    required this.rememberMe,
    required this.onRememberMeChanged,
    required this.onForgotPassword,
    required this.onLogin,
  });

  @override
  State<LoginTab> createState() => _LoginTabState();
}

class _LoginTabState extends State<LoginTab> {
  bool _obscurePassword = true;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        ValidatedTextField(
          icon: Icons.person,
          hint: "Email",
          controller: widget.emailController,
          validator: Validators.validateEmail,
        ),
        const VSpace.md(),
        ValidatedTextField(
          icon: Icons.lock,
          hint: "Password",
          obscure: _obscurePassword,
          controller: widget.passwordController,
          validator: Validators.validatePassword,
          toggleObscure: () {
            setState(() {
              _obscurePassword = !_obscurePassword;
            });
          },
        ),
        const VSpace.sm(),
        RememberMeCheckbox(
          value: widget.rememberMe,
          onChanged: widget.onRememberMeChanged,
          onForgotPassword: widget.onForgotPassword,
        ),
        const VSpace.md(),
        SubmitButton(
          text: "Log in",
          onPressed: widget.onLogin,
          isLoading: false,
        ),
      ],
    );
  }
}
