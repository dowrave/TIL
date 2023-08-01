using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Background : MonoBehaviour
{
    // 배경 머티리얼 속성 정의
    public Material bgMaterial;
    public float scrollSpeed = 0.2f;


    // Start is called before the first frame update
    void Start()
    {
        
    }

    // 살아 있는 동안
    void Update()
    {
        // 윗방향으로
        Vector2 direction = Vector2.up;
        
        // 이렇게 스크롤 할래요
        bgMaterial.mainTextureOffset += direction * scrollSpeed * Time.deltaTime;
    }
}
