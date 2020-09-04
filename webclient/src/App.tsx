import React from 'react';

import {
  View, Text,
} from 'react-native';

export type AppProps = {
};

export type AppState = {
}

export default class App extends React.Component<AppProps, AppState> {
  constructor(props: AppProps) {
    super(props);
  }

  render(): JSX.Element {
    return (
      <View>
        <Text>{'Number 7'}</Text>
      </View>
    );
  }
}
