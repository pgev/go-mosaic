import React from 'react';

import {
  Button, Dimensions, StyleSheet, View, Text, TextInput,
} from 'react-native';

import { gql, request } from 'graphql-request';

export type AppProps = {
  endpoint: string;
};

export type AppState = {
  value: string;
}

interface TData {
  value: string;
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

    this.readValue = this.readValue.bind(this);
    this.updateValue = this.updateValue.bind(this);

    this.state = {
      value: '',
    };
  }

  onChangeText(value: string): void {
    this.setState({
      value,
    });
  }

  onUpdate(): void {
    this.updateValue(this.state.value);
  }

  readValue(): void {
    const query = gql`
      {
        value
      }
    `;

    request(
      this.props.endpoint,
      query,
    )
      .then((data: TData) => {
        this.setState({
          value: data.value,
        });
      });
  }

  updateValue(value: string): void {
    const query = gql`
      mutation UpdateValue{$value: String!) {
        updateValue(value: $value) {
          value
        }
      }
    `;

    const variables = {
      value,
    };

    request(
      this.props.endpoint,
      query,
      variables,
    )
      .then((data: TData) => {
        this.setState({
          value: data.value,
        });
      });
  }

  componentDidMount(): void {
    this.readValue();
  }

  render(): JSX.Element {
    return (
      <View style={styles.topContainer}>
        <Text>{'Store Your Secret Number'}</Text>
        <View style={{ flexDirection: 'row', padding: 2 }}>
          <Text>{'Your Secret Number:'}</Text>
          <TextInput
            keyboardType = 'numeric'
            onChangeText={this.onChangeText}
            maxLength={2}
            value = {this.state.value}
            style={{ borderWidth: 1 }}
          />
        </View>
        <Button title='Update' onPress={this.onUpdate} />
      </View>
    );
  }
}
