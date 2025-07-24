import 'package:flutter/material.dart';
import 'package:frontend/screens/pages/views/project_create.dart';
import 'package:frontend/models/project_model.dart';
import 'package:frontend/utils/token_storage.dart';
import 'package:frontend/services/api_service.dart';
import 'package:frontend/widgets/discover/widgets.dart';

typedef ProjectTapCallback = void Function(Project project);

class MyProjectsPage extends StatefulWidget {
  const MyProjectsPage({super.key, this.onProjectTap});
  final ProjectTapCallback? onProjectTap;

  @override
  State<MyProjectsPage> createState() => _MyProjectsPageState();
}

class _MyProjectsPageState extends State<MyProjectsPage> {
  List<Project> _projects = [];
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _fetchProjects();
  }

  Future<void> _fetchProjects() async {
    setState(() => _loading = true);
    try {
      final token = await getToken();
      if (token == null) throw Exception('No token');
      final projects = await ApiService.getUserProjects(token);
      setState(() {
        _projects = projects.map((e) => Project.fromJson(e)).toList();
        _loading = false;
      });
    } catch (e) {
      setState(() => _loading = false);
      if (!mounted) return;
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(SnackBar(content: Text('Ошибка загрузки проектов: $e')));
    }
  }

  void _onProjectCreated() {
    _fetchProjects();
    ScaffoldMessenger.of(
      context,
    ).showSnackBar(const SnackBar(content: Text('Проект успешно создан!')));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : _projects.isEmpty
          ? const Center(child: Text('Нет проектов'))
          : ListView.builder(
              itemCount: _projects.length,
              itemBuilder: (context, i) => ProjectCard(
                project: _projects[i],
                onTap: () => widget.onProjectTap?.call(_projects[i]),
              ),
            ),
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.add),
        onPressed: () async {
          await Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) =>
                  ProjectCreate(onProjectCreated: _onProjectCreated),
            ),
          );
        },
      ),
    );
  }
}
