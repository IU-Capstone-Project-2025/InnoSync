import 'package:flutter/material.dart';
import '../../theme/dimensions.dart';

/// Vertical spacing widget for consistent spacing between elements
class VSpace extends StatelessWidget {
  final double height;

  const VSpace(this.height, {super.key});

  // Predefined spacing sizes
  const VSpace.xs({super.key}) : height = AppDimensions.xs;
  const VSpace.sm({super.key}) : height = AppDimensions.sm;
  const VSpace.md({super.key}) : height = AppDimensions.md;
  const VSpace.lg({super.key}) : height = AppDimensions.lg;
  const VSpace.xl({super.key}) : height = AppDimensions.xl;
  const VSpace.xxl({super.key}) : height = AppDimensions.xxl;
  const VSpace.xxxl({super.key}) : height = AppDimensions.xxxl;

  // Legacy spacing for backward compatibility
  const VSpace.small({super.key}) : height = AppDimensions.sm;
  const VSpace.medium({super.key}) : height = AppDimensions.md;
  const VSpace.large({super.key}) : height = AppDimensions.xl;
  const VSpace.mediumMinus({super.key}) : height = AppDimensions.md - 2;
  const VSpace.mediumPlus({super.key}) : height = AppDimensions.md + 4;
  const VSpace.mediumPlusPlus({super.key}) : height = AppDimensions.md + 8;

  @override
  Widget build(BuildContext context) {
    return SizedBox(height: height);
  }
}

/// Horizontal spacing widget for consistent spacing between elements
class HSpace extends StatelessWidget {
  final double width;

  const HSpace(this.width, {super.key});

  // Predefined spacing sizes
  const HSpace.xs({super.key}) : width = AppDimensions.xs;
  const HSpace.sm({super.key}) : width = AppDimensions.sm;
  const HSpace.md({super.key}) : width = AppDimensions.md;
  const HSpace.lg({super.key}) : width = AppDimensions.lg;
  const HSpace.xl({super.key}) : width = AppDimensions.xl;
  const HSpace.xxl({super.key}) : width = AppDimensions.xxl;
  const HSpace.xxxl({super.key}) : width = AppDimensions.xxxl;

  // Legacy spacing for backward compatibility
  const HSpace.small({super.key}) : width = AppDimensions.sm;
  const HSpace.medium({super.key}) : width = AppDimensions.md;
  const HSpace.large({super.key}) : width = AppDimensions.xl;
  const HSpace.mediumPlus({super.key}) : width = AppDimensions.md + 8;

  @override
  Widget build(BuildContext context) {
    return SizedBox(width: width);
  }
}

/// Fixed height container for layout purposes
class FixedHeightSpace extends StatelessWidget {
  final double height;
  final Widget child;

  const FixedHeightSpace({
    super.key,
    required this.height,
    required this.child,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(height: height, child: child);
  }
}

/// Fixed width container for layout purposes
class FixedWidthSpace extends StatelessWidget {
  final double width;
  final Widget child;

  const FixedWidthSpace({super.key, required this.width, required this.child});

  @override
  Widget build(BuildContext context) {
    return SizedBox(width: width, child: child);
  }
}

/// Flexible spacing that takes available space
class FlexibleSpace extends StatelessWidget {
  const FlexibleSpace({super.key});

  @override
  Widget build(BuildContext context) {
    return const Spacer();
  }
}
