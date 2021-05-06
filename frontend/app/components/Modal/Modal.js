export default class Modal extends React.PureComponent {
	constructor(props) {
    super(props);
    this.el = document.createElement('div');
  }

	render() {
    return ReactDOM.createPortal(
      this.props.children,
      this.el,
    );
	}
}