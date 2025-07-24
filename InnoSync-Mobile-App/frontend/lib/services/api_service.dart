import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart' as http;

class ApiService {
  static const String baseUrl = 'http://localhost:3000/api/v1';

  static Future<String?> login(String email, String password) async {
    final url = Uri.parse('$baseUrl/auth/login');
    final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'email': email, 'password': password}),
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return data['tokens']?['access_token'];
    }
    return null;
  }

  static Future<bool> signup(
    String email,
    String password,
    String fullName,
  ) async {
    final url = Uri.parse('$baseUrl/auth/signup');
    final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'email': email,
        'password': password,
        'full_name': fullName,
      }),
    );
    return response.statusCode == 201;
  }

  static Future<List<Map<String, dynamic>>> getUserProjects(
    String token,
  ) async {
    final url = Uri.parse('$baseUrl/projects/me');
    final response = await http.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );
    if (response.statusCode == 200) {
      final decoded = jsonDecode(response.body);
      final List data = decoded is List ? decoded : [];
      return data.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to load projects');
    }
  }

  static Future<bool> createProject(
    String token,
    String title,
    String? description,
    int? teamSize,
  ) async {
    final url = Uri.parse('$baseUrl/projects');
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'title': title,
        'description': description,
        'team_size': teamSize,
      }),
    );
    return response.statusCode == 201;
  }

  Future<List<Map<String, dynamic>>> fetchInvitations(String token) async {
    final url = Uri.parse('$baseUrl/invitations/received');
    final response = await http.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );
    if (response.statusCode == 200) {
      final decoded = jsonDecode(response.body);
      final List data = decoded is List ? decoded : [];
      return data.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to load invitations (status: ${response.statusCode}, body: ${response.body})');
    }
  }

  Future<List<Map<String, dynamic>>> fetchProposals(String token) async {
    final url = Uri.parse('$baseUrl/applications');
    final response = await http.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );
    if (response.statusCode == 200) {
      final decoded = jsonDecode(response.body);
      final List data = decoded is List ? decoded : [];
      return data.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to load proposals (status: ${response.statusCode}, body: ${response.body})');
    }
  }

  Future<List<Map<String, dynamic>>> fetchChats(String token) async {
    final url = Uri.parse('$baseUrl/chats'); // Проверь правильный эндпоинт
    final response = await http.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );
    if (response.statusCode == 200) {
      final decoded = jsonDecode(response.body);
      final List data = decoded is List ? decoded : [];
      return data.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to load chats');
    }
  }

  Future<Map<String, dynamic>> fetchProfile(String token) async {
    final url = Uri.parse('$baseUrl/profile/me');
    final response = await http.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );
    if (response.statusCode == 200) {
      return jsonDecode(response.body) as Map<String, dynamic>;
    } else {
      throw Exception('Failed to load profile');
    }
  }

  Future<bool> createInvitation({
    required String token,
    required int projectRoleId,
    required int recipientId,
    String? message,
  }) async {
    final url = Uri.parse('$baseUrl/invitations');
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'project_role_id': projectRoleId,
        'recipient_id': recipientId,
        if (message != null) 'message': message,
      }),
    );
    if (response.statusCode == 201) {
      return true;
    } else {
      return false;
    }
  }

  Future<bool> createProfile({
    required String token,
    required String name,
    required String email,
    String? telegram,
    String? github,
    String? bio,
    String? position,
    String? education,
    String? expertise,
    String? expertiseLevel,
    List<String>? technologies,
    List<Map<String, dynamic>>? workExperience,
    String? resumeUrl,
    File? profileImage,
  }) async {
    final url = Uri.parse('$baseUrl/profile');
    final Map<String, dynamic> body = {'name': name, 'email': email};
    if (telegram != null && telegram.trim().isNotEmpty) {
      body['telegram'] = telegram;
    }
    if (github != null && github.trim().isNotEmpty) body['github'] = github;
    if (bio != null && bio.trim().isNotEmpty) body['bio'] = bio;
    if (position != null && position.trim().isNotEmpty) {
      body['position'] = position;
    }
    if (education != null && education.trim().isNotEmpty) {
      body['education'] = education;
    }
    if (expertise != null && expertise.trim().isNotEmpty) {
      body['expertise'] = expertise;
    }
    if (expertiseLevel != null && expertiseLevel.trim().isNotEmpty) {
      body['expertise_level'] = expertiseLevel;
    }
    if (technologies != null && technologies.isNotEmpty) {
      body['technologies'] = technologies;
    }
    if (workExperience != null && workExperience.isNotEmpty) {
      body['work_experience'] = workExperience;
    }
    if (resumeUrl != null && resumeUrl.isNotEmpty) {
      body['resume_url'] = resumeUrl;
    }
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(body),
    );
    return response.statusCode == 201;
  }

  Future<bool> applyToProject({
    required String token,
    required int projectId,
    String? message,
  }) async {
    final url = Uri.parse('$baseUrl/applications');
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'project_id': projectId,
        if (message != null) 'message': message,
      }),
    );
    return response.statusCode == 201;
  }

  Future<bool> respondToInvitation({
    required String token,
    required int invitationId,
    required String response, // 'ACCEPTED' или 'DECLINED'
  }) async {
    final url = Uri.parse('$baseUrl/invitations/$invitationId/respond');
    final resp = await http.patch(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({'response': response}),
    );
    return resp.statusCode == 200;
  }

  Future<bool> createProjectRole({
    required String token,
    required int projectId,
    required String roleName,
    String? description,
    List<String>? skills,
    int? count,
  }) async {
    final url = Uri.parse('$baseUrl/projects/$projectId/roles');
    final body = {
      'role_name': roleName,
      if (description != null && description.isNotEmpty) 'description': description,
      if (skills != null && skills.isNotEmpty) 'skills': skills,
      if (count != null) 'count': count,
    };
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(body),
    );
    return response.statusCode == 201;
  }
}
