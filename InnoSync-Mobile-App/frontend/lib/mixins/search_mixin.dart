import 'package:flutter/material.dart';

mixin SearchMixin<T extends StatefulWidget> on State<T> {
  final TextEditingController searchController = TextEditingController();
  final List<String> searchFilters = <String>[];

  void addSearchFilter(String value) {
    final keyword = value.trim();
    if (keyword.isNotEmpty && !searchFilters.contains(keyword)) {
      setState(() {
        searchFilters.add(keyword);
      });
    }
    searchController.clear();
  }

  void removeSearchFilter(String keyword) {
    setState(() {
      searchFilters.remove(keyword);
    });
  }

  void clearAllFilters() {
    setState(() {
      searchFilters.clear();
    });
  }

  List<String> getFilteredItems(List<String> items) {
    if (searchFilters.isEmpty) return items;

    return items.where((item) {
      return searchFilters.any(
        (filter) => item.toLowerCase().contains(filter.toLowerCase()),
      );
    }).toList();
  }

  @override
  void dispose() {
    searchController.dispose();
    super.dispose();
  }
}
