import 'package:flutter/material.dart';
import 'blocs/bloc.dart';
import 'blocs/provider.dart';

class BlocApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Provider(
      child: MaterialApp(
        title: 'Log me in',
        home: Scaffold(
          body: BlocLoginScreen(),
        ),
      ),
    );
  }
}

class BlocLoginScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final bloc = Provider.of(context);

    return Container(
      margin: const EdgeInsets.all(20.0),
      child: Column(
        children: [
          emailField(bloc),
          passwordField(bloc),
          Container(margin: const EdgeInsets.only(top: 25.0)),
          submitButton(bloc),
        ],
      ),
    );
  }
}

Widget emailField(bloc) {
  return StreamBuilder<String>(
    stream: bloc.email,
    builder: (context, snapshot) {
      return TextField(
        onChanged: bloc.changeEmail,
        keyboardType: TextInputType.emailAddress,
        decoration: InputDecoration(
          hintText: 'you@example.com',
          labelText: 'Email',
          errorText: snapshot.error == null ? null : snapshot.error as String,
        ),
      );
    },
  );
}

Widget passwordField(bloc) {
  return StreamBuilder(
      stream: bloc.password,
      builder: (context, snapshot) {
        return TextField(
          onChanged: bloc.changePassword,
          obscureText: true,
          decoration: InputDecoration(
            hintText: 'Password',
            labelText: 'Password',
            errorText: snapshot.error == null ? null : snapshot.error as String,
          ),
        );
      });
}

Widget submitButton(bloc) {
  return StreamBuilder(
    stream: bloc.submitValid,
    builder: (context, snapshot) {
      return ElevatedButton(
        style: ElevatedButton.styleFrom(backgroundColor: Colors.blue),
        onPressed: !snapshot.hasData ? null : bloc.submit,
        child: const Text('Login'),
      );
    },
  );
}
