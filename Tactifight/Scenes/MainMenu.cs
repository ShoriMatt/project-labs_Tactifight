using Godot;

public partial class MainMenu : Control
{
	private Button _buttonCommencer;
	private Button _buttonOptions;
	private Button _buttonQuitter;

	public override void _Ready()
	{
		// Récupère les boutons par leur nom dans la scène
		_buttonCommencer = GetNode<Button>("CenterContainer/VBoxContainer/Button_Commencer");
		_buttonOptions   = GetNode<Button>("CenterContainer/VBoxContainer/Button_Options");
		_buttonQuitter   = GetNode<Button>("CenterContainer/VBoxContainer/Button_Quitter");

		// Connecte les signaux "Pressed" des boutons à des fonctions
		_buttonCommencer.Pressed += OnCommencerPressed;
		_buttonOptions.Pressed   += OnOptionsPressed;
		_buttonQuitter.Pressed   += OnQuitterPressed;
	}

	private void OnCommencerPressed()
	{
		GD.Print("Démarrage du jeu !");
		// Tu pourras plus tard charger la scène principale :
		// GetTree().ChangeSceneToFile("res://Scenes/Battle.tscn");
	}

	private void OnOptionsPressed()
	{
		GD.Print("Menu Options à venir...");
	}

	private void OnQuitterPressed()
	{
		GD.Print("Quitter le jeu !");
		GetTree().Quit();
	}
}
