using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Background : MonoBehaviour
{
    // ��� ��Ƽ���� �Ӽ� ����
    public Material bgMaterial;
    public float scrollSpeed = 0.2f;


    // Start is called before the first frame update
    void Start()
    {
        
    }

    // ��� �ִ� ����
    void Update()
    {
        // ����������
        Vector2 direction = Vector2.up;
        
        // �̷��� ��ũ�� �ҷ���
        bgMaterial.mainTextureOffset += direction * scrollSpeed * Time.deltaTime;
    }
}
