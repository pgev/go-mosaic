import React from 'react';

import {
  Button, Dimensions, StyleSheet, View, Text, TextInput,
} from 'react-native';

import { gql, request } from 'graphql-request';

export type AppProps = {
  endpoint: string;
};

export type AppState = {
  secret: string;
}

const styles = StyleSheet.create({
  topContainer: {
    flex: 1,
    flexDirection: 'column',
    alignItems: 'stretch',
    width: Math.round(Dimensions.get('window').width),
    height: Math.round(Dimensions.get('window').height),
    padding: 2,
    borderWidth: 1,
  },
});

export default class App extends React.Component<AppProps, AppState> {
  constructor(props: AppProps) {
    super(props);

    this.onChangeText = this.onChangeText.bind(this);
    this.onUpdate = this.onUpdate.bind(this);

    this.state = {
      secret: '',
    };
  }

  onChangeText(secret: string): void {
    this.setState({
      secret,
    });
  }

  onUpdate(): void {
    this.updateSecret(this.state.secret);
  }

  readSecret(): void {
    const query = gql`
      {
        Secret {
          Value
        }
      }
    `;
    request(
      this.props.endpoint,
      query,
    )
      .then((data: any) => {
        console.log('Reading ...');
        console.log(data);
        this.setState({
          secret: data.Secret.Value,
        });
      })
      .catch((reason: any) => {
        console.error(reason);
      });
  }

  updateSecret(secret: string): void {
    const query = gql`
      mutation($secret: String!) {
        UpdateSecretValue(Secret: $secret) {
          Value
        }
      }
    `;

    const variables = {
      secret,
    };

    request(
      this.props.endpoint,
      query,
      variables,
    )
      .then((data: any) => {
        this.setState({
          secret: data.UpdateSecretValue.Value,
        });
      })
      .catch((reason: any) => {
        console.error(reason);
      });
  }

  componentDidMount(): void {
    this.readSecret();
  }

  render(): JSX.Element {
    return (
      <View style={styles.topContainer}>
        <Text>{'Store Your Secret'}</Text>
        <View style={{ flexDirection: 'row', padding: 2 }}>
          <Text>{'Your Secret:'}</Text>
          <TextInput
            onChangeText={this.onChangeText}
            maxLength={100}
            value = {this.state.secret}
            style={{ borderWidth: 1 }}
          />
        </View>
        <Button title='Update' onPress={this.onUpdate} />
      </View>
    );
  }
}
