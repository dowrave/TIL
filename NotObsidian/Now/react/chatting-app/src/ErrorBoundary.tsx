import { Component, ErrorInfo } from 'react'

type Props = { children: JSX.Element };

type State = { hasError: boolean;
            errorMessage: string };

export default class ErrorBoundary extends Component<Props, State> {
  state = { hasError: false, errorMessage: ""};

  static getDerivedStateFromError(error: Error): State {
    // 에러 발생 시 hasError 상태를 true로 전환한다.
    return { hasError: true, errorMessage: error.message }
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // 에러 로깅, 서버로의 전송 등을 수행한다.
    console.log("에러 발생 ! ");
    console.log("에러 명 : ", error.name);
    console.log("에러 메시지 : ", error.message);
    console.log("컴포넌트 스택 : ", errorInfo.componentStack);;
  }

  render() {
        // 에러 발생 시 fallback UI 렌더링
        if (this.state.hasError) {
            return (
                <div>
                <h2>에러 발생</h2>
                <hr />
                <p>에러 메시지 : {this.state.errorMessage} </p>
                </div>
            );
        }
    return this.props.children;

  }
}