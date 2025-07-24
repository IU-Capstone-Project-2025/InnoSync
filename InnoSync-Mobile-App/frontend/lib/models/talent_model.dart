class Talent {
  final String name;
  final String description;
  final String? profileImageUrl;
  final int yearsOfExperience;
  final String graduationType;
  final List<String> skills;
  final String? location;
  final String? email;
  final List<String> positions;

  const Talent({
    required this.name,
    required this.description,
    this.profileImageUrl,
    required this.yearsOfExperience,
    required this.graduationType,
    this.skills = const [],
    this.positions = const [],
    this.location,
    this.email,
  });
}
