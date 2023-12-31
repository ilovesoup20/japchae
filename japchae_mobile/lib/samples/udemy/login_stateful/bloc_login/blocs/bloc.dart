import 'dart:async';
import 'package:japchae_mobile/samples/udemy/login_stateful/bloc_login/blocs/validators.dart';
import 'package:rxdart/rxdart.dart';

// class Bloc extends Validators {
class Bloc extends Object with Validators {
  final _email = BehaviorSubject<String>();
  final _password = BehaviorSubject<String>();

  // Add data to stream
  Function(String) get changeEmail => _email.sink.add;
  Function(String) get changePassword => _password.sink.add;

  // Retrieve data from stream
  Stream<String> get email => _email.stream.transform(validateEmail);
  Stream<String> get password => _password.stream.transform(validatePassword);
  Stream<bool> get submitValid =>
      Rx.combineLatest2(email, password, (e, p) => true);

  submit() {
    final validEmail = _email.value;
    final validPassword = _password.value;

    print('Email: $validEmail, PW: $validPassword');
  }

  dispose() {
    _email.close();
    _password.close();
  }
}

// Global instance
// final bloc = Bloc();
