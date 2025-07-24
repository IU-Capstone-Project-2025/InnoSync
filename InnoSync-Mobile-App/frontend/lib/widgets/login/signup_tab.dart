import 'package:flutter/material.dart';
import 'validated_text_field.dart';
import '../../widgets/common/spacing.dart';
import '../../widgets/common/submit_button.dart';
import '../../utils/validators.dart';

class SignupTab extends StatefulWidget {
  final TextEditingController emailController;
  final TextEditingController passwordController;
  final VoidCallback onSignup;

  const SignupTab({
    super.key,
    required this.emailController,
    required this.passwordController,
    required this.onSignup,
  });

  @override
  State<SignupTab> createState() => _SignupTabState();
}

class _SignupTabState extends State<SignupTab> {
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
        const VSpace.medium(),
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
        const VSpace.mediumPlus(),
        SubmitButton(
          text: "Continue",
          onPressed: widget.onSignup,
          isLoading: false,
        ),
      ],
    );
  }
}
