using Godot;

public partial class Option : Control
{
	private Button _returnButton;

	public override void _Ready()
	{
		_returnButton = GetNode<Button>("Button");

		// Connexion correcte du signal
		_returnButton.Pressed += OnReturnPressed;
	}

	private void OnReturnPressed()
	{
		GD.Print("Retour au menu");
		GetTree().ChangeSceneToFile("res://Scenes/SpaceMenu_Demo.tscn");
	}
}
