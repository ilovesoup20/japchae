import 'package:flutter/material.dart';
import '../mixins/validation_mixin.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<StatefulWidget> createState() {
    return LoginScreenState();
  }
}

class LoginScreenState extends State<LoginScreen> with ValidationMixin {
  final formKey = GlobalKey<FormState>();
  String? email = '';
  String? password = '';

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(20.0),
      child: Form(
        key: formKey,
        child: Column(children: [
          emailField(),
          Container(margin: const EdgeInsets.only(top: 25.0)),
          passwordField(),
          Container(margin: const EdgeInsets.only(top: 25.0)),
          submitButton(),
        ]),
      ),
    );
  }

  Widget emailField() {
    return TextFormField(
      keyboardType: TextInputType.emailAddress,
      decoration: const InputDecoration(
        labelText: 'Email Address',
        hintText: 'you@example.com',
      ),
      validator: validateEmail,
      onSaved: (String? value) {
        print(value);
        email = value;
      },
    );
  }

  Widget passwordField() {
    return TextFormField(
        obscureText: true,
        decoration: const InputDecoration(
          labelText: 'Password',
          hintText: 'Password',
        ),
        validator: validatePassword,
        onSaved: (value) {
          print(value);
          password = value;
        });
  }

  Widget submitButton() {
    return ElevatedButton(
      onPressed: () {
        // formKey.currentState?.reset();
        var formState = formKey.currentState;
        if (formState != null && formState.validate()) {
          formKey.currentState?.save();
          print('$email and $password');
        }
      },
      child: const Text('Submit'),
    );
  }
}
