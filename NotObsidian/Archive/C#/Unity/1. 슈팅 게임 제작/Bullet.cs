using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Bullet : MonoBehaviour
{
    public float speed = 5;

    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        // 1. 방향을 구한다
        Vector3 dir = Vector3.up;
        // 2. 이동한다
        transform.position += dir * speed * Time.deltaTime;
    }
}
