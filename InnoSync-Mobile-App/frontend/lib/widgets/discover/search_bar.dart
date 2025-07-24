import 'package:flutter/material.dart';

class CustomSearchBar extends StatelessWidget {
  final TextEditingController controller;
  final Function(String) onSubmitted;
  final Function(String) onAddFilter;
  final String hintText;

  const CustomSearchBar({
    super.key,
    required this.controller,
    required this.onSubmitted,
    required this.onAddFilter,
    this.hintText = 'Search...',
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Expanded(
      child: TextField(
        controller: controller,
        onSubmitted: onSubmitted,
        decoration: InputDecoration(
          hintText: hintText,
          prefixIcon: const Icon(Icons.search),
          border: OutlineInputBorder(borderRadius: BorderRadius.circular(12.0)),
          filled: true,
          fillColor: theme.colorScheme.surface,
          contentPadding: const EdgeInsets.symmetric(
            vertical: 0,
            horizontal: 16,
          ),
          suffixIcon: IconButton(
            icon: const Icon(Icons.add),
            onPressed: () => onAddFilter(controller.text),
          ),
        ),
      ),
    );
  }
}
