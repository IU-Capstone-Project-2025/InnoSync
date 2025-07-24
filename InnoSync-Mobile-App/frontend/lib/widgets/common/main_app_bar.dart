import 'package:flutter/material.dart';

class MainAppBar extends StatelessWidget implements PreferredSizeWidget {
  final String? title;
  final List<Widget>? actions;
  final bool showBackButton;

  const MainAppBar({
    super.key,
    this.title,
    this.actions,
    this.showBackButton = true,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return AppBar(
      leading: showBackButton
          ? IconButton(
              icon: const Icon(Icons.arrow_back),
              onPressed: () => Navigator.of(
                context,
              ).pushNamedAndRemoveUntil('/main', (route) => false),
            )
          : null,
      backgroundColor: theme.scaffoldBackgroundColor,
      foregroundColor: theme.colorScheme.onSurface,
      elevation: 0,
      title: title != null
          ? Text(
              title!,
              style: theme.textTheme.titleLarge?.copyWith(fontSize: 20),
            )
          : null,
      centerTitle: true,
      actions: actions,
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}
