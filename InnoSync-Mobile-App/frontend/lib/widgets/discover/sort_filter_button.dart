import 'package:flutter/material.dart';
import 'package:frontend/widgets/common/spacing.dart';

class SortFilterButton extends StatelessWidget {
  final VoidCallback? onTap;
  final String? label;

  const SortFilterButton({super.key, this.onTap, this.label});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Material(
      color: theme.colorScheme.surface,
      borderRadius: BorderRadius.circular(12.0),
      child: InkWell(
        borderRadius: BorderRadius.circular(12.0),
        onTap: onTap,
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 10),
          child: Row(
            children: [
              Icon(Icons.sort, color: theme.colorScheme.onSurface),
              if (label != null) ...[
                const HSpace.small(),
                Text(
                  label!,
                  style: theme.textTheme.bodyMedium?.copyWith(
                    color: theme.colorScheme.onSurface,
                  ),
                ),
              ],
              const HSpace.small(),
              Icon(
                Icons.arrow_drop_down,
                color: theme.colorScheme.onSurface,
                size: 20,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
