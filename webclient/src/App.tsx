import React from 'react';

import {
  Dimensions, StyleSheet, View, Text, TextInput,
} from 'react-native';

export type AppProps = {
};

export type AppState = {
  number?: string;
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

    this.state = {
      number: '',
    };
  }

  onChangeText(number: string): void {
    this.setState({
      number,
    });
  }

  render(): JSX.Element {
    return (
      <View style={styles.topContainer}>
        <Text>{'Number 7 App'}</Text>
        <View style={{flexDirection: 'row', padding: 2}}>
          <Text>{'Number:'}</Text>
          <TextInput
            placeholder={'Please, update number ...'}
            keyboardType = 'numeric'
            onChangeText={this.onChangeText}
            maxLength={2}
            value = {this.state.number}
          />
        </View>
      </View>
    );
  }
}
