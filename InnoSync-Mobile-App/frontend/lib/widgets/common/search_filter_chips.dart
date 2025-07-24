import 'package:flutter/material.dart';

class SearchFilterChips extends StatelessWidget {
  final List<String> filters;
  final Function(String) onRemoveFilter;

  const SearchFilterChips({
    super.key,
    required this.filters,
    required this.onRemoveFilter,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    if (filters.isEmpty) return const SizedBox.shrink();

    return Column(
      children: [
        const SizedBox(height: 12),
        SizedBox(
          height: 40,
          child: ListView.separated(
            scrollDirection: Axis.horizontal,
            itemCount: filters.length,
            separatorBuilder: (context, i) => const SizedBox(width: 8),
            itemBuilder: (context, i) {
              final keyword = filters[i];
              return Chip(
                label: Text(
                  keyword,
                  style: theme.textTheme.bodyMedium?.copyWith(
                    color: theme.colorScheme.onPrimary,
                  ),
                ),
                backgroundColor: theme.colorScheme.primary,
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(24),
                ),
                deleteIcon: Container(
                  decoration: BoxDecoration(
                    color: theme.colorScheme.onPrimary,
                    shape: BoxShape.circle,
                  ),
                  child: Icon(
                    Icons.close,
                    size: 18,
                    color: theme.colorScheme.primary,
                  ),
                ),
                onDeleted: () => onRemoveFilter(keyword),
              );
            },
          ),
        ),
      ],
    );
  }
}
