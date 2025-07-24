import 'package:flutter/material.dart';
import '../../theme/dimensions.dart';

class TabSelector extends StatelessWidget {
  final TabController controller;
  final List<String> tabLabels;

  const TabSelector({
    super.key,
    required this.controller,
    required this.tabLabels,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Container(
      height: AppDimensions.tabHeight,
      decoration: BoxDecoration(
        border: Border.all(color: theme.colorScheme.primary, width: 2),
        borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
        color: theme.colorScheme.surface,
      ),
      child: TabBar(
        controller: controller,
        tabs: tabLabels.map((label) => Tab(text: label)).toList(),
        indicator: BoxDecoration(
          color: theme.colorScheme.primary,
          borderRadius: BorderRadius.circular(AppDimensions.radiusMd),
          border: Border.all(color: theme.colorScheme.primary, width: 2),
        ),
        labelColor: theme.colorScheme.onPrimary,
        unselectedLabelColor: theme.colorScheme.primary,
        indicatorSize: TabBarIndicatorSize.tab,
        dividerColor: Colors.transparent,
        indicatorPadding: EdgeInsets.zero,
      ),
    );
  }
}
