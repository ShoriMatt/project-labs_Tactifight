shader_type canvas_item;

/// Position UV 0–1 de la planète dans l’écran
uniform vec2  center = vec2(0.92, 0.14);
/// Rayon (0–1 par rapport à l’écran)
uniform float radius = 0.22;
/// Couleur de base et liseré
uniform vec3  base = vec3(0.90, 0.50, 0.20);
uniform vec3  rim  = vec3(1.00, 0.85, 0.55);

void fragment() {
    vec2 p = UV - center;
    float r = length(p);

    // masque disque (alpha)
    float disk = smoothstep(radius, radius - 0.005, r);
    if (disk <= 0.0) { COLOR = vec4(0.0);
	}

    // bord lumineux (atmosphère)
    float edge = smoothstep(radius * 0.90, radius, r);
    vec3 col = mix(base, rim, pow(edge, 2.0));

    COLOR = vec4(col, disk);
}
