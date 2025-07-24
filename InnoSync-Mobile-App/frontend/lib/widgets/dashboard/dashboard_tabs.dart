import 'package:flutter/material.dart';

class DashboardTabs extends StatefulWidget {
  final List<String> tabLabels;
  final List<Widget> tabViews;
  final Color? indicatorColor;

  const DashboardTabs({
    super.key,
    required this.tabLabels,
    required this.tabViews,
    this.indicatorColor,
  }) : assert(tabLabels.length == tabViews.length);

  @override
  State<DashboardTabs> createState() => _DashboardTabsState();
}

class _DashboardTabsState extends State<DashboardTabs>
    with TickerProviderStateMixin {
  late TabController _tabController;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(
      length: widget.tabLabels.length,
      vsync: this,
    );
  }

  @override
  void dispose() {
    _tabController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Column(
      children: [
        TabBar(
          controller: _tabController,
          indicatorColor: widget.indicatorColor ?? theme.colorScheme.primary,
          tabs: widget.tabLabels.map((label) => Tab(text: label)).toList(),
        ),
        Expanded(
          child: TabBarView(
            controller: _tabController,
            children: widget.tabViews,
          ),
        ),
      ],
    );
  }
}
