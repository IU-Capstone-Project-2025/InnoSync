import 'package:flutter/material.dart';
import 'switcher_button.dart';

class ModeSwitcher extends StatelessWidget {
  final String currentMode;
  final Function(String) onModeChanged;
  final List<String> modes;

  const ModeSwitcher({
    super.key,
    required this.currentMode,
    required this.onModeChanged,
    this.modes = const ['Projects', 'Talents'],
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Container(
          decoration: BoxDecoration(
            color: theme.colorScheme.onSurface.withValues(alpha: 0.1),
            borderRadius: BorderRadius.circular(32),
          ),
          padding: const EdgeInsets.all(4),
          child: Row(
            children: modes
                .map(
                  (mode) => SwitcherButton(
                    label: mode,
                    selected: currentMode == mode,
                    onTap: () => onModeChanged(mode),
                  ),
                )
                .toList(),
          ),
        ),
      ],
    );
  }
}
