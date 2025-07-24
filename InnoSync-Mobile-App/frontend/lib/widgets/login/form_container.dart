import 'package:flutter/material.dart';
import '../../theme/dimensions.dart';
import '../common/widgets.dart';

class FormContainer extends StatelessWidget {
  final Widget child;
  final String? title;
  final String? subtitle;
  final GlobalKey<FormState> formKey;

  const FormContainer({
    super.key,
    required this.child,
    required this.formKey,
    this.title,
    this.subtitle,
  });

  @override
  Widget build(BuildContext context) {
    final colorScheme = Theme.of(context).colorScheme;
    final textTheme = Theme.of(context).textTheme;
    return Center(
      child: SingleChildScrollView(
        child: Card(
          color: colorScheme.surface,
          elevation: 4,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(16),
          ),
          child: Container(
            width: AppDimensions.containerWidth,
            padding: const EdgeInsets.symmetric(
              horizontal: AppDimensions.paddingXl,
              vertical: AppDimensions.paddingXl,
            ),
            child: Form(
              key: formKey,
              child: Column(
                children: [
                  if (title != null) ...[
                    Text(title!, style: textTheme.headlineSmall?.copyWith(fontWeight: FontWeight.bold)),
                    const VSpace.xs(),
                  ],
                  if (subtitle != null) ...[
                    Text(subtitle!, style: textTheme.bodyMedium),
                    const VSpace.xl(),
                  ],
                  child,
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
